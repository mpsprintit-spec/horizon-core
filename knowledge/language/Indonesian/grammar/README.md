# Grammar Module

## Overview

This README defines the documentation foundation for Indonesian grammatical categories, morphology, inflectional behavior, and rule explanations. It establishes how future Indonesian language knowledge should be organized while preserving Horizon's principle that knowledge supports responsible, context-aware assistance.

## Purpose

The purpose of this document is to provide standards for future grammar documentation without writing grammar datasets. It helps contributors add future resources consistently, review changes predictably, and keep Indonesian language knowledge aligned with Horizon's human-first and knowledge-centered design.

## Scope

This module includes documentation standards, classification rules, metadata expectations, and small illustrative examples for Indonesian grammatical categories, morphology, inflectional behavior, and rule explanations.

This module does not include bulk data, generated datasets, implementation code, AI prompts, business logic, and material owned by another module. Contributors must not use this README to populate large data collections or change repository architecture.

## Relationship

This module interacts with syntax, semantics, dictionary, and examples. Cross-references should clarify ownership: each concept should be defined in one primary module and referenced from other modules only when needed. Shared terminology must remain consistent across the Indonesian Language Knowledge Base.

## Directory Structure

The intended directory layout is:

```text
knowledge/language/Indonesian/grammar/
└── README.md
```

Do not create additional directories until there is a clear documentation need. New files should reuse the current module structure and should not duplicate content that already belongs elsewhere.

## File Naming Convention

Use lowercase, predictable, human-readable file names. Prefer hyphen-separated words, such as `usage-notes.md` or `politeness-levels.md`. Avoid spaces, mixed casing, dates as primary identifiers, and ambiguous abbreviations. File names should describe the documented concept, not the author, implementation status, or temporary task.

## Writing Standard

Write documentation in professional English. Use Markdown headings in sentence case, short paragraphs, and bullet lists for rules or checklists. Use Indonesian terms only when they are the subject of documentation, and provide concise English explanations on first use. Keep terminology consistent with the other Indonesian language modules, especially the terms `module`, `entry`, `usage`, `meaning`, `context`, `example`, and `cross-reference`.

## Quality Standard

Acceptable documentation must be accurate, concise, reviewable, and maintainable. Each document should state its scope, avoid duplicate definitions, and identify related modules when concepts overlap. Reviewers should verify structural consistency, terminology consistency, Markdown formatting, and alignment with The Book of Horizon before approving changes.

## Examples

- Document word classes, affix behavior, reduplication, particles, tense-aspect marking, and other grammatical patterns.
- Use rules and constraints rather than long lists of generated forms.
- Reference syntax when grammar affects phrase or clause structure.

Examples in this module are documentation examples only. They are not authoritative dataset entries and should not be treated as complete linguistic coverage.

## Future Expansion

This module may grow by adding focused Markdown files that document one stable concept at a time. Expansion should preserve the current architecture, use cross-references instead of duplication, and prioritize knowledge that improves safe, contextual understanding for human assistance.
