import * as React from "react";
import axios from "axios";
import "../App.scss";

import {
  addVotedIDToStorage,
  isVotedIdInStorage,
} from "../interfaces/votedCookieInterface";

import BallotInterface from "../interfaces/ballotInterface";
import VoteInterface from "../interfaces/voteInterface";
import promptInterface from "../interfaces/promptInterface";
import { BallotContext } from "../context/BallotContext";

export default function BallotComponent() {
  const { currentBallot, setCurrentBallot } = React.useContext(BallotContext);
  
  // does the post request for the actual vote
  const DoVote = (p: promptInterface) => {
    const vote: VoteInterface = { ballotID: currentBallot.id, promptID: p.id };
    axios
    .post<BallotInterface>("http://139.144.18.143:8080/vote", vote)
    .then((response) => {
      setCurrentBallot(response.data);
    })
    .catch((ex) => {
      const error = "An unexpected error has occurred. Could not vote";
    });
    addVotedIDToStorage(currentBallot.id);
  };
  
  return (
    <div>
    {/* <h1 className="heading">Vote for this</h1> */}
    <h2>{currentBallot.description}</h2>
    {/* <h3>Prompts are:</h3> */}
    <ul>
    {currentBallot.prompts.map((p) => (
      <div className="panel">
      <li key={p.id}>
      <p>{p.name}</p>
      <p>{p.description}</p>
      <p>{p.votes}</p>
      {
        // only render the button if the ballot id is not in the list of already voted ids
        //!isVotedIdInStorage(currentBallot.id) && (
        <button onClick={() => DoVote(p)} className="button is-green">Vote</button>
        // <Button onClick={() => DoVote(p)} as="a" variant="primary">
        // Vote
        // </Button>
        //)
      }
      </li>
      </div>
      ))}
      </ul>
      </div>
      );
    }
    