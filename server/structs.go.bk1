package main

// promptStruct represents data about a record promptStruct.
type promptStruct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Votes       int64  `json:"votes"`
}

type ballotStruct struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Prompts     []promptStruct `json:"prompts"`
}

type voteStruct struct {
	BallotID string `json:"ballotID"`
	PromptID string `json:"promptID"`
}

var ballots = []ballotStruct{
	{ID: "1", Description: "Welcome to DANCExDANCE. There will be moments in the performance when you will be invited to participate via this webpage.", Prompts: []promptStruct{
	}},
	{ID: "2", Description: "What is cuter?", Prompts: []promptStruct{
		{ID: "1", Name: "Puppies", Description: "", Votes: 0},
		{ID: "2", Name: "Kittens", Description: "", Votes: 0},
		{ID: "3", Name: "Babies", Description: "", Votes: 0},
	}},
	{ID: "3", Description: "Which color do you prefer?", Prompts: []promptStruct{
		{ID: "1", Name: "Red", Description: "", Votes: 0},
		{ID: "2", Name: "Green", Description: "", Votes: 0},
		{ID: "3", Name: "Blue", Description: "", Votes: 0},
	}},
	{ID: "4", Description: "Do you feel like your smartphone knows you better than you know yourself?", Prompts: []promptStruct{
		{ID: "1", Name: "Yes", Description: "", Votes: 0},
		{ID: "2", Name: "No", Description: "", Votes: 0},
	}},
	{ID: "5", Description: "Should we get on with the show already?", Prompts: []promptStruct{
		{ID: "1", Name: "Yes", Description: "", Votes: 0},
		{ID: "2", Name: "No", Description: "", Votes: 0},
	}},
	{ID: "6", Description: "Do you believe that humans and robots will live in harmony?", Prompts: []promptStruct{
		{ID: "1", Name: "Yes", Description: "", Votes: 0},
		{ID: "2", Name: "No", Description: "", Votes: 0},
	}},
	{ID: "7", Description: "Choice 1", Prompts: []promptStruct{
		{ID: "1", Name: "Turn Light Green", Description: "", Votes: 0},
		{ID: "2", Name: "Turn Light Red", Description: "", Votes: 0},
		{ID: "3", Name: "Turn Light Blue", Description: "", Votes: 0},
	}},
	{ID: "8", Description: "Choice 2", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
	}},
	{ID: "9", Description: "Choice 3", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Move to knee", Description: "", Votes: 0},
	}},
	{ID: "10", Description: "Choice 4", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Move back and forth", Description: "", Votes: 0},
	}},
	{ID: "11", Description: "Choice 5", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Blink lights 5 times", Description: "", Votes: 0},
	}},
	{ID: "12", Description: "Choice 6", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Move to wrist", Description: "", Votes: 0},
	}},
	{ID: "13", Description: "Choice 7", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Move to ankle", Description: "", Votes: 0},
	}},
	{ID: "14", Description: "Choice 8", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Move to belly", Description: "", Votes: 0},
	}},
	{ID: "15", Description: "Choice 9", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Move to ankle", Description: "", Votes: 0},
	}},
	{ID: "16", Description: "Choice 10", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Vibrate", Description: "", Votes: 0},
	}},
	{ID: "17", Description: "Choice 11", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Move to shoulder", Description: "", Votes: 0},
	}},
	{ID: "18", Description: "Choice 12", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Stop and start to belly", Description: "", Votes: 0},
	}},
	{ID: "19", Description: "Choice 13", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Move to shoulder via the ankle", Description: "", Votes: 0},
	}},
	{ID: "20", Description: "Choice 14", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Cycle through colors and vibrate", Description: "", Votes: 0},
	}},
	{ID: "21", Description: "Choice 15", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Move to knee and then to shoulder", Description: "", Votes: 0},
	}},
	{ID: "22", Description: "Choice 16", Prompts: []promptStruct{
		{ID: "1", Name: "Continue Choreography", Description: "", Votes: 0},
		{ID: "2", Name: "Leave dancer via the wrist", Description: "", Votes: 0},
	}},
	{ID: "23", Description: "Choice 17", Prompts: []promptStruct{
		{ID: "1", Name: "Turn Light Green", Description: "", Votes: 0},
		{ID: "2", Name: "Turn Light Red", Description: "", Votes: 0},
		{ID: "3", Name: "Turn Light Blue", Description: "", Votes: 0},
	}},
}
