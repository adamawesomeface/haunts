package actions

import (
  "encoding/gob"
  "haunts/base"
  "haunts/game/status"
  "haunts/game"
  "path/filepath"
)

func registerAttacks() map[string]func() game.Action {
  attack_actions := make(map[string]*BasicAttackDef)
  base.RemoveRegistry("actions-attack_actions")
  base.RegisterRegistry("actions-attack_actions", attack_actions)
  base.RegisterAllObjectsInDir("actions-attack_actions", filepath.Join(base.GetDataDir(), "actions", "basic_attacks"), ".json", "json")
  makers := make(map[string]func() game.Action)
  for name := range attack_actions {
    cname := name
    makers[cname] = func() game.Action {
      a := BasicAttack{ Defname: cname }
      base.GetObject("actions-attack_actions", &a)
      return &a
    }
  }
  return makers
}

func init() {
  game.RegisterActionMakers(registerAttacks)
  gob.Register(&BasicAttack{})
}

// Basic Attacks are single target and instant, they are also readyable
type BasicAttack struct {
  Defname string
  *BasicAttackDef
}
type BasicAttackDef struct {
  Name     string
  Kind     status.Kind
  Ap       int
  Strength int
  Range    int
}

func (a *BasicAttack) Readyable() bool {
  return true
}
func (a *BasicAttack) Cost() int {
  return a.Ap
}
func (a *BasicAttack) Prep(*game.Entity) bool {
  return true
}
func (a *BasicAttack) HandleInput() bool {
  return true
}
func (a *BasicAttack) HandleOutput() {
}
func (a *BasicAttack) Cancel() {
}
func (a *BasicAttack) Maintain(dt int64) game.MaintenanceStatus {
  // Do stuff
  return game.Complete
}
func (a *BasicAttack) Interrupt() bool {
  return true
}

