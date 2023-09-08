import axios from "axios";
import React from "react";
import _ from "lodash";
import BallotComponent from "../components/BallotComponent";
import { BallotContext } from "../context/BallotContext";
import BallotInterface from "../interfaces/ballotInterface";
import GlobalState from "../interfaces/globalState";
import { StateContext } from "../context/StateContext";

export const client = axios.create({
  baseURL: "http://172.234.43.240:8080",
});

export default function Voting() {
  const defaultBallot: BallotInterface = {
    id: "",
    description: "Hello welcome to DANCExDANCE",
    prompts: [],
  };
  
  const defaultState: GlobalState = {
    fade: false,
  };
  
  const [currentBallot, _setCurrentBallot]: [BallotInterface, any] =
  React.useState(defaultBallot);
  
  const [currentState, _setGlobalState]: [GlobalState, any] =
  React.useState(defaultState);
  
  // This confusing syntax exists to ensure that the setcurrentballot only actually gets called when the new polled ballot is different from the existing current ballot
  // if we keep re-writing the current ballot, the BallotComponent will keep re-rendering. Not a big issue in this app, but this re-rendering is not good for performance in general
  // see: https://stackoverflow.com/a/70712897/9533616
  const setCurrentBallot = (bal: BallotInterface) => {
    _setCurrentBallot((prev) => {
      if (!_.isEqual(bal, prev)) {
        return bal;
      }
      return prev;
    });
  };
  
  const setGlobalState = (gstate: GlobalState) => {
    _setGlobalState((prev) => {
      if (!_.isEqual(gstate, prev)) {
        console.log("Global state changed");
        return gstate;
      }
      return prev;
    });
  };
  
  // error and loading states
  const [loading, setLoading]: [boolean, (loading: boolean) => void] =
  React.useState<boolean>(true);
  const [error, setError]: [string, (error: string) => void] =
  React.useState("");
  
  // Polls the new current ballot. Split into it's own function because we use the same code twice: the very first poll and the regular interval polls
  function getCurrentBallot() {
    client
    .get<BallotInterface>("/ballots/current", { timeout: 10000 })
    .then((response) => {
      setCurrentBallot(response.data);
      setLoading(false);
    })
    .catch((ex) => {
      const error = " "
    //   ex.code === "ECONNABORTED"
    //   ? "A timeout has occurred"
    //   : ex.response.status === 404
    //   ? "Resource Not found"
    //   : "An unexpected error has occurred";
      setError(error);
      setLoading(false);
    });
  }
  
  function getGlobalState() {
    client
    .get<GlobalState>("/mysecretkey/fadestate", { timeout: 10000 })
    .then((response) => {
      setGlobalState(response.data);
      setLoading(false);
    })
    .catch((ex) => {
      const error = " "
    //   ex.code === "ECONNABORTED"
    //   ? "A timeout has occurred"
    //   : ex.response.status === 404
    //   ? "Resource Not found"
    //   : "An unexpected error has occurred";
      setError(error);
      setLoading(false);
    });
  }
  
  
  React.useEffect(() => {
    // poll the current ballot for the first time
    getCurrentBallot();
    getGlobalState();
    
    // run timer at regular intervals to get new current ballot
    const timer = setInterval(() => {
      getCurrentBallot();
      getGlobalState();
    }
    , 100);
    return () => clearInterval(timer);
  }, []);
  
  
  
  return (
    <StateContext.Provider value={{ currentState, setGlobalState}}>
    <BallotContext.Provider value={{ currentBallot, setCurrentBallot}}>
    <BallotComponent />
    {error && <p className="error">{error}</p>}
    {loading && <p>Loading</p>}
    </BallotContext.Provider>
    </StateContext.Provider>
    );
  }