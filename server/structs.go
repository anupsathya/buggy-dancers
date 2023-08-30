package main

// promptStruct represents data about a record promptStruct.
type promptStruct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Votes       int64  `json:"votes"`
}

type ballotStruct struct {
	ID          string         `json:"id"`
	Description string         `json:"description"`
	Prompts     []promptStruct `json:"prompts"`
}

type voteStruct struct {
	BallotID string `json:"ballotID"`
	PromptID string `json:"promptID"`
}

type fadeState struct {
	Fade bool `json:"fade"`
}

var defaultFade = fadeState{true}

var ballots = []ballotStruct{
	{ID: "1", Description: "Welcome to DANCExDANCE. There will be moments in the performance when you will be invited to participate via this webpage.  Please make sure to set your phone screens to stay 'awake' for the duration of the performance", Prompts: []promptStruct{}},
	{ID: "2", Description: "Which is cuter?", Prompts: []promptStruct{
		{ID: "1", Name: "Puppies", Description: "", Votes: 0},
		{ID: "2", Name: "Kittens", Description: "", Votes: 0},
		{ID: "3", Name: "Babies", Description: "", Votes: 0},
	}},
	{ID: "3", Description: "Which color do you prefer?", Prompts: []promptStruct{
		{ID: "1", Name: "Red", Description: "", Votes: 0},
		{ID: "2", Name: "Green", Description: "", Votes: 0},
		{ID: "3", Name: "Blue", Description: "", Votes: 0},
	}},
	{ID: "4", Description: "Do you control your technology or does your technology control you?", Prompts: []promptStruct{
		{ID: "1", Name: "Technology", Description: "", Votes: 0},
		{ID: "2", Name: "You", Description: "", Votes: 0},
	}},
	{ID: "5", Description: "Should we get on with the show already?", Prompts: []promptStruct{
		{ID: "1", Name: "Yes", Description: "", Votes: 0},
		{ID: "2", Name: "No", Description: "", Votes: 0},
	}},
	{ID: "6", Description: "Do you believe that humans and robots will live in harmony?", Prompts: []promptStruct{
		{ID: "1", Name: "Yes", Description: "", Votes: 0},
		{ID: "2", Name: "No", Description: "", Votes: 0},
	}},
	{ID: "7", Description: "Turn the robot's light...", Prompts: []promptStruct{
		{ID: "1", Name: "Red", Description: "", Votes: 0},
		{ID: "2", Name: "Green", Description: "", Votes: 0},
		{ID: "3", Name: "Blue", Description: "", Votes: 0},
	}},
	{ID: "8", Description: "Should the robot...", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography (move to right arm)", Description: "", Votes: 0},
		{ID: "2", Name: "Blinks", Description: "", Votes: 0},
	}},
	{ID: "9", Description: "Should the robot...", Prompts: []promptStruct{
		{ID: "1", Name: "Continue choreography (travel across and all the way to the belly)", Description: "", Votes: 0},
		{ID: "2", Name: "Reverse to shoulder (in opposite direction)", Description: "", Votes: 0},
	}},
	{ID: "10", Description: "Should the robot...", Prompts: []promptStruct{
		{ID: "1", Name: "Continue choreography (change color)", Description: "", Votes: 0},
		{ID: "2", Name: "Move back and forth", Description: "", Votes: 0},
	}},
	{ID: "11", Description: "Should the robot...", Prompts: []promptStruct{
		{ID: "1", Name: "Continue choreography (travel to left arm)", Description: "", Votes: 0},
		{ID: "2", Name: "Backtrack (to knee)", Description: "", Votes: 0},
	}},
	{ID: "12", Description: "Should the robot...", Prompts: []promptStruct{
		{ID: "1", Name: "Continue choreography (travel to ankle)", Description: "", Votes: 0},
		{ID: "2", Name: "Move back and forth", Description: "", Votes: 0},
	}},
	{ID: "13", Description: "Should the robot...", Prompts: []promptStruct{
		{ID: "1", Name: "Continue choreography (exit at the ankle)", Description: "", Votes: 0},
		{ID: "2", Name: "Vibrate", Description: "", Votes: 0},
	}},
	{ID: "14", Description: "How do you feel about Artificial Intelligence technology?", Prompts: []promptStruct{
		{ID: "1", Name: "Curious", Description: "", Votes: 0},
		{ID: "2", Name: "Ambivalent", Description: "", Votes: 0},
		{ID: "3", Name: "Fearful", Description: "", Votes: 0},
	}},
	{ID: "15", Description: "Which best describes how you encounter the unknown?", Prompts: []promptStruct{
		{ID: "1", Name: "Uncertainty", Description: "", Votes: 0},
		{ID: "2", Name: "Excitement", Description: "", Votes: 0},
	}},
	{ID: "16", Description: "Does technology do more to connect us or disconnect us from each other?", Prompts: []promptStruct{
		{ID: "1", Name: "Connect", Description: "", Votes: 0},
		{ID: "2", Name: "Disconnect", Description: "", Votes: 0},
	}},
	{ID: "17", Description: "What color best symbolizes the future?", Prompts: []promptStruct{
		{ID: "1", Name: "Red", Description: "", Votes: 0},
		{ID: "2", Name: "Green", Description: "", Votes: 0},
		{ID: "3", Name: "Blue", Description: "", Votes: 0},
	}},
	{ID: "18", Description: "Which word would you like to see used in a poem? ", Prompts: []promptStruct{
		{ID: "1", Name: "Technology", Description: "", Votes: 0},
		{ID: "2", Name: "Unknown", Description: "", Votes: 0},
		{ID: "3", Name: "Connect", Description: "", Votes: 0},
		{ID: "4", Name: "Color", Description: "", Votes: 0},
	}},
	{ID: "19", Description: "", Prompts: []promptStruct{}},
	// {ID: "20", Description: "Choice 14", Prompts: []promptStruct{
	// 	{ID: "1", Name: "Continue Choreography (move to wrist)", Description: "", Votes: 0},
	// 	{ID: "2", Name: "Cycle through colors and vibrate", Description: "", Votes: 0},
	// }},
	// {ID: "21", Description: "Choice 15", Prompts: []promptStruct{
	// 	{ID: "1", Name: "Continue Choreography (move down)", Description: "", Votes: 0},
	// 	{ID: "2", Name: "Move to knee and then to shoulder", Description: "", Votes: 0},
	// }},
	// {ID: "22", Description: "Choice 16", Prompts: []promptStruct{
	// 	{ID: "1", Name: "Continue Choreography (exit via ankle)", Description: "", Votes: 0},
	// 	{ID: "2", Name: "Leave dancer via the wrist", Description: "", Votes: 0},
	// }},
	// {ID: "23", Description: "Choice 17", Prompts: []promptStruct{
	// 	{ID: "1", Name: "Turn Light Green", Description: "", Votes: 0},
	// 	{ID: "2", Name: "Turn Light Red", Description: "", Votes: 0},
	// 	{ID: "3", Name: "Turn Light Blue", Description: "", Votes: 0},
	// }},
}
