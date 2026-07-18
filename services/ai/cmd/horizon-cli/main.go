package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/project-horizon/horizon-core/services/ai/core"
	"github.com/project-horizon/horizon-core/services/ai/perception"
	"github.com/project-horizon/horizon-core/services/ai/plugin"
	"github.com/project-horizon/horizon-core/services/ai/thinking"
	"github.com/project-horizon/horizon-core/services/ai/websearch"
)

const memoryFile = "brain_memory.json"

type cli struct {
	horizon    *core.HorizonEngine
	perceiver  perception.UserInputPerception
	web        *websearch.Engine
	in         *bufio.Scanner
	out        io.Writer
	debug      bool
	history    []string
	context    []string
	last       pipelineResult
	startedAt  time.Time
	memoryPath string
}

type pipelineResult struct {
	Input         string
	Tokens        []string
	Answer        string
	Confidence    float64
	Activation    []string
	Understanding string
	Reasoning     string
	Context       []string
	Hypothesis    []string
	NeedsWeb      bool
	Learned       bool
}

func main() {
	app := newCLI(os.Stdin, os.Stdout, memoryFile)
	if err := app.startup(); err != nil {
		fmt.Fprintf(app.out, "Startup warning: %v\n", err)
	}
	app.run(context.Background())
}

func newCLI(input io.Reader, output io.Writer, memoryPath string) *cli {
	horizon := core.NewHorizonEngine()
	horizon.Execution.RegisterPlugin("terbang", &plugin.DronePlugin{})
	horizon.Execution.RegisterPlugin("logsystem", &plugin.ChatbotPlugin{})
	return &cli{
		horizon:    horizon,
		web:        websearch.NewEngine(nil),
		in:         bufio.NewScanner(input),
		out:        output,
		startedAt:  time.Now(),
		memoryPath: memoryPath,
	}
}

func (c *cli) startup() error {
	c.printBanner()
	err := c.horizon.Knowledge.Load(c.memoryPath)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(c.out, "! Memory load skipped: %v\n", err)
	}
	fmt.Fprintln(c.out, "✓ Neural Memory")
	fmt.Fprintln(c.out, "✓ Activation Engine")
	fmt.Fprintln(c.out, "✓ Understanding Engine")
	fmt.Fprintln(c.out, "✓ Reasoning Engine")
	fmt.Fprintln(c.out, "✓ Language Engine")
	fmt.Fprintln(c.out, "✓ Learning Engine")
	fmt.Fprintln(c.out, "✓ WebSearch Engine")
	if err != nil && os.IsNotExist(err) {
		return nil
	}
	return err
}

func (c *cli) printBanner() {
	fmt.Fprintln(c.out, "======================================")
	fmt.Fprintln(c.out, " Horizon Cognitive Intelligence")
	fmt.Fprintln(c.out, " Neural Semantic Memory Ready")
	fmt.Fprintln(c.out, "======================================")
}

func (c *cli) run(ctx context.Context) {
	for {
		fmt.Fprint(c.out, "> ")
		if !c.in.Scan() {
			break
		}
		input := strings.TrimSpace(c.in.Text())
		if input == "" {
			continue
		}
		if c.handleCommand(input) {
			if input == "exit" || input == "quit" {
				break
			}
			continue
		}
		result := c.process(ctx, input)
		c.last = result
		c.history = append(c.history, input)
		if len(c.history) > 25 {
			c.history = c.history[len(c.history)-25:]
		}
		if c.debug {
			c.printDebug(result)
		}
		fmt.Fprintln(c.out, result.Answer)
	}
	if err := c.horizon.Knowledge.Save(c.memoryPath); err != nil {
		fmt.Fprintf(c.out, "Shutdown save warning: %v\n", err)
	}
	c.printSessionStats()
}

func (c *cli) process(ctx context.Context, input string) pipelineResult {
	signals, err := c.perceiver.Perceive(input)
	if err != nil || len(signals) == 0 {
		return pipelineResult{Input: input, Answer: "Input tidak dapat dipersepsi."}
	}
	tokens := signals[0].Tokens
	thought, ok := c.horizon.Thinking.Think(input, c.context)
	if !ok && !isQuestion(input) {
		c.horizon.Learning.Assimilate(input, 0.8)
		thought, ok = c.horizon.Thinking.Think(input, c.context)
	}
	learned := false
	if !isQuestion(input) {
		c.horizon.Learning.Assimilate(input, 0.85)
		learned = true
	}
	if c.web.ShouldSearch(thought.Confidence, len(c.horizon.Thinking.LastState.UnknownNodes), len(thought.Conflicts)) {
		thought.NeedsWebSearch = true
		_, _ = c.web.Perceive(ctx, input)
	}
	answer := c.horizon.Language.Generate(thought)
	if !ok && learned {
		answer = "Informasi baru dipelajari dan Neural Memory diperbarui."
	}
	c.context = deriveContext(tokens, thought.Concepts)
	return pipelineResult{
		Input:         input,
		Tokens:        tokens,
		Answer:        answer,
		Confidence:    thought.Confidence,
		Activation:    rankedTokens(c.horizon),
		Understanding: strings.Join(thought.Concepts, ", "),
		Reasoning:     answer,
		Context:       c.context,
		Hypothesis:    hypothesisSummary(thought),
		NeedsWeb:      thought.NeedsWebSearch,
		Learned:       learned,
	}
}

func (c *cli) handleCommand(input string) bool {
	switch strings.ToLower(input) {
	case "help":
		fmt.Fprintln(c.out, "Commands: help, exit, quit, memory, nodes, synapses, activation, context, history, clear, save, load, debug on, debug off")
	case "exit", "quit":
		fmt.Fprintln(c.out, "Shutting down Horizon...")
	case "memory":
		fmt.Fprintf(c.out, "Memory: %d nodes, %d synapses\n", c.nodeCount(), c.synapseCount())
	case "nodes":
		fmt.Fprintf(c.out, "Nodes: %d\n", c.nodeCount())
	case "synapses":
		fmt.Fprintf(c.out, "Synapses: %d\n", c.synapseCount())
	case "activation":
		fmt.Fprintf(c.out, "Activation: %s\n", strings.Join(c.last.Activation, ", "))
	case "context":
		fmt.Fprintf(c.out, "Context: %s\n", strings.Join(c.context, ", "))
	case "history":
		fmt.Fprintln(c.out, strings.Join(c.history, "\n"))
	case "clear":
		fmt.Fprint(c.out, "\033[H\033[2J")
	case "save":
		if err := c.horizon.Knowledge.Save(c.memoryPath); err != nil {
			fmt.Fprintf(c.out, "Save failed: %v\n", err)
		} else {
			fmt.Fprintln(c.out, "Memory saved.")
		}
	case "load":
		if err := c.horizon.Knowledge.Load(c.memoryPath); err != nil {
			fmt.Fprintf(c.out, "Load failed: %v\n", err)
		} else {
			fmt.Fprintln(c.out, "Memory loaded.")
		}
	case "debug on":
		c.debug = true
		fmt.Fprintln(c.out, "Debug enabled.")
	case "debug off":
		c.debug = false
		fmt.Fprintln(c.out, "Debug disabled.")
	default:
		return false
	}
	return true
}

func (c *cli) printDebug(r pipelineResult) {
	fmt.Fprintf(c.out, "[debug] Activation: %s\n", strings.Join(r.Activation, ", "))
	fmt.Fprintf(c.out, "[debug] Understanding: %s\n", r.Understanding)
	fmt.Fprintf(c.out, "[debug] Reasoning: %s\n", r.Reasoning)
	fmt.Fprintf(c.out, "[debug] Confidence: %.2f\n", r.Confidence)
	fmt.Fprintf(c.out, "[debug] Context: %s\n", strings.Join(r.Context, ", "))
	fmt.Fprintf(c.out, "[debug] Hypothesis: %s\n", strings.Join(r.Hypothesis, ", "))
	fmt.Fprintf(c.out, "[debug] WebSearch: %t\n", r.NeedsWeb)
	fmt.Fprintf(c.out, "[debug] Learning: %t\n", r.Learned)
}

func (c *cli) printSessionStats() {
	fmt.Fprintf(c.out, "Session: %d inputs, %d nodes, %d synapses, %s uptime\n", len(c.history), c.nodeCount(), c.synapseCount(), time.Since(c.startedAt).Round(time.Millisecond))
}
func (c *cli) nodeCount() int { return len(c.horizon.Knowledge.Registry.Nodes()) }
func (c *cli) synapseCount() int {
	total := 0
	for _, n := range c.horizon.Knowledge.Registry.Nodes() {
		total += len(n.Synapses)
	}
	return total
}

func isQuestion(input string) bool {
	s := strings.ToLower(strings.TrimSpace(input))
	return strings.HasSuffix(s, "?") ||
		strings.HasPrefix(s, "apa ") ||
		strings.HasPrefix(s, "siapa ") ||
		strings.HasPrefix(s, "mengapa ") ||
		strings.HasPrefix(s, "kenapa ") ||
		strings.HasPrefix(s, "bagaimana ")
}
func deriveContext(tokens, concepts []string) []string {
	out := append([]string{}, concepts...)
	if len(out) == 0 {
		out = append(out, tokens...)
	}
	if len(out) > 5 {
		return out[:5]
	}
	return out
}
func rankedTokens(h *core.HorizonEngine) []string {
	nodes := h.Thinking.LastState.ActivationHistory
	if len(nodes) == 0 {
		return nil
	}
	ranked := nodes[len(nodes)-1].RankedNodes
	out := make([]string, 0, len(ranked))
	for _, n := range ranked {
		out = append(out, n.Token)
	}
	sort.Strings(out)
	return out
}
func hypothesisSummary(thought thinking.Thought) []string {
	out := make([]string, 0, len(thought.Hypotheses))
	for _, h := range thought.Hypotheses {
		out = append(out, fmt.Sprintf("confidence=%.2f evidence=%d conflicts=%d", h.Confidence, h.Evidence, h.Conflicts))
	}
	return out
}
