// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Client-side methods for editing a match in the match review page.

const scoreTemplate = Handlebars.compile($("#scoreTemplate").html());
const allianceResults = {};
let matchResult;

// Hijack the form submission to inject the data in JSON form so that it's easier for the server to parse.
$("form").submit(function() {
  updateResults("red");
  updateResults("blue");

  matchResult.RedScore = allianceResults["red"].score;
  matchResult.BlueScore = allianceResults["blue"].score;
  matchResult.RedCards = allianceResults["red"].cards;
  matchResult.BlueCards = allianceResults["blue"].cards;
  const matchResultJson = JSON.stringify(matchResult);

  // Inject the JSON data into the form as hidden inputs.
  $("<input />").attr("type", "hidden").attr("name", "matchResultJson").attr("value", matchResultJson).appendTo("form");

  return true;
});

// Draws the match-editing form for one alliance based on the cached result data.
const renderResults = function(alliance) { // TIGER_TODO
  const result = allianceResults[alliance];
  const scoreContent = scoreTemplate(result);
  $("#" + alliance + "Score").html(scoreContent);

  // Set the values of the form fields from the JSON results data.
  for (let i = 0; i < 1; i++) {
    const i1 = i + 1;

    getInputElement(alliance, "EndgameStatuses" + i1, result.score.EndgameStatuses[i]).prop("checked", true);

  }

  getInputElement(alliance, "ConesHigh").val(result.score.HighCones);
  getInputElement(alliance, "ConesLow").val(result.score.LowCones);

  if (result.score.Fouls != null) {
    $.each(result.score.Fouls, function(k, v) {
      getInputElement(alliance, "Foul" + k + "IsTechnical").prop("checked", v.IsTechnical);
      getInputElement(alliance, "Foul" + k + "Team", v.TeamId).prop("checked", true);
      getSelectElement(alliance, "Foul" + k + "RuleId").val(v.RuleId);
    });
  }

  if (result.cards != null) {
    $.each(result.cards, function(k, v) {
      getInputElement(alliance, "Team" + k + "Card", v).prop("checked", true);
    });
  }
};

// Converts the current form values back into JSON structures and caches them.
const updateResults = function(alliance) { // TIGER_TODO
  const result = allianceResults[alliance];
  const formData = {};
  $.each($("form").serializeArray(), function(k, v) {
    formData[v.name] = v.value;
  });

  result.score.HighCones = parseInt(formData[alliance + "ConesHigh"])
  result.score.LowCones = parseInt(formData[alliance + "ConesLow"])

  result.score.EndgameStatuses = [];
  for (let i = 0; i < 1; i++) {
    const i1 = i + 1;

    result.score.EndgameStatuses[i] = parseInt(formData[alliance + "EndgameStatuses" + i1]);
  }

  result.score.Fouls = [];

  for (let i = 0; formData[alliance + "Foul" + i + "Index"]; i++) {
    const prefix = alliance + "Foul" + i;
    const foul = {
      IsTechnical: formData[prefix + "IsTechnical"] === "on",
      TeamId: parseInt(formData[prefix + "Team"]),
      RuleId: parseInt(formData[prefix + "RuleId"]),
    };
    result.score.Fouls.push(foul);
  }

  result.cards = {};
  $.each([result.team1], function(i, team) {
    result.cards[team] = formData[alliance + "Team" + team + "Card"];
  });
};

// Appends a blank foul to the end of the list.
const addFoul = function(alliance) {
  updateResults(alliance);
  const result = allianceResults[alliance];
  result.score.Fouls.push({IsTechnical: false, TeamId: 0, Rule: 0});
  renderResults(alliance);
};

// Removes the given foul from the list.
const deleteFoul = function(alliance, index) {
  updateResults(alliance);
  const result = allianceResults[alliance];
  result.score.Fouls.splice(index, 1);
  renderResults(alliance);
};

// Returns the form input element having the given parameters.
const getInputElement = function(alliance, name, value) {
  let selector = "input[name=" + alliance + name + "]";
  if (value !== undefined) {
    selector += "[value=" + value + "]";
  }
  return $(selector);
};

// Returns the form select element having the given parameters.
const getSelectElement = function(alliance, name) {
  const selector = "select[name=" + alliance + name + "]";
  return $(selector);
};
