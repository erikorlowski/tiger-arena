// Copyright 2020 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model of a game-specific rule.

package game

type Rule struct {
	Id             int
	RuleNumber     string
	IsTechnical    bool
	IsRankingPoint bool
	Description    string
}

// All rules from the 2022 game that carry point penalties.
var rules = []*Rule{ // TIGER_TODO
	{1, "G104", false, false, "Robots may not have any parts extend more than 18 inches from their frame."},
	{2, "G201", false, false, "Strategies clearly aimed at forcing the opponent ALLIANCE to violate a rule are not in the spirit of FIRST Robotics Competition and not allowed."},
	{3, "G201", true, false, "Strategies clearly aimed at forcing the opponent ALLIANCE to violate a rule are not in the spirit of FIRST Robotics Competition and not allowed. If REPEATED, TECH FOUL."},
	{4, "G202", false, false, "ROBOTS may not PIN an opponent’s ROBOT for more than 5 seconds."},
	{5, "G202", false, false, "A Robot may not use a Component outside its frame to initiate contact with an opponent Robot inside the vertical projection of that opponent Robot’s frame."},
	{6, "G203", true, false, " A Robot may not damage an opponent Robot, deliberately, or by using a Component outside its frame to initiate contact with an opponent Robot inside the vertical projection of that opponent Robot’s frame."},
	{7, "G204", true, false, "A Robot may not deliberately tip or entangle an opposing Robot."},
	{8, "G205", false, false, "Robots may not touch an opposing Robot that is In their Loading Zone or Low Goal. (i.e. If the blue Robot is In the blue Loading Zone, the red Robot may not touch them)"},
	{9, "G206", true, false, "Robots may not be In the opposing Loading Zone."},
	{10, "G301", false, false, "Robots may not control more than one Cone at a time. Control is considered to be carrying a Cone, or pushing a Cone across the field for a strategic purpose beyond moving around the Field."},
	{11, "G301", true, false, "Robots may not control more than one Cone at a time. Control is considered to be carrying a Cone, or pushing a Cone across the field for a strategic purpose beyond moving around the Field. Tech Foul for every Cone scored while in violation of this rule."},
	{12, "G302", true, false, "Robots may not descore Cones from the opposing team’s Low Goal or High Goal."},
	{13, "H201", false, false, "Team members must stay in their designated space behind their end of the Field during matches. Foul if repeated during a match."},
	{14, "H202", false, false, "Human Players must only place Cones In their Loading Zone and may not interact with any Cones that are not In their Loading Zone."},
	{15, "H203", false, false, "A Human Player may never be In the field when a Robot is In their Loading Zone. A Human Player may never touch a Cone that is touching a Robot."},
}
var ruleMap map[int]*Rule

// Returns the rule having the given ID, or nil if no such rule exists.
func GetRuleById(id int) *Rule {
	return GetAllRules()[id]
}

// Returns a slice of all defined rules that carry point penalties.
func GetAllRules() map[int]*Rule {
	if ruleMap == nil {
		ruleMap = make(map[int]*Rule, len(rules))
		for _, rule := range rules {
			ruleMap[rule.Id] = rule
		}
	}
	return ruleMap
}
