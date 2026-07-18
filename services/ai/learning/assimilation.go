package learning

import (
	"fmt"
	"github.com/project-horizon/horizon-core/services/ai/knowledge"
)

func (l *LearningUnit) Assimilate(subjekStr, relasiStr, objekStr string, bobot float64) {
	subjek := l.Kb.Store(subjekStr)
	relasi := l.Kb.Store(relasiStr)
	objek := l.Kb.Store(objekStr)

	if !l.WebSearchSimulate(subjekStr, objekStr) {
		l.Kb.Mu.Lock()
		subjek.Inhibitions[relasi] = true
		l.Kb.Mu.Unlock()
		fmt.Printf("🔍 [LEARNING] Kontradiksi dideteksi! Saraf penghambat aktif untuk '%s' -> [%s]\n", subjekStr, objekStr)
		return
	}

	l.Kb.Mu.Lock()
	subjek.Relations[relasi] = &knowledge.Synapse{TargetNode: objek, Weight: bobot}
	l.Kb.Mu.Unlock()
}
