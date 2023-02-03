import React, { createContext } from "react";
import GlobalState from "../interfaces/globalState";

const defaultBallot: GlobalState = {
  fade: false,
};
// const [currentBallot, setCurrentBallot]: [BallotInterface, (ballots: BallotInterface) => void] = React.useState(defaultBallot);

export const StateContext = createContext<any>({});
