import React from "react";
import axios from "axios";
import BallotComponent from "../components/BallotComponent";
import { BallotContext } from "../context/BallotContext";
import BallotInterface from "../interfaces/ballotInterface";

export default function Dashboard() {
  const [ballots, setBallots] = React.useState<any[]>([]);
  const fetchUsers = async () => {
    const response = await axios.get("http://172.234.43.240:8080/mysecretkey/ballots");
    setBallots(response.data);
    console.log(response.data[0].id);
    console.log("test log");
  };
  
  React.useEffect(() => {
    fetchUsers();
  }, []);
  
  const cueThis = (p) => {
    axios
    .get("http://172.234.43.240:8080/mysecretkey/ballots/setCurrent/" + p.id)
    .then((response) => {
      console.log(response.data);
    })
    .catch((ex) => {
      const error = "An unexpected error has occurred. Could not vote";
    });        
  };

  const toggleFade = () => {
    axios
    .get("http://172.234.43.240:8080/mysecretkey/togglefadestate")
    .then((response) => {
      console.log(response.data);
    })
    .catch((ex) => {
      const error = "An unexpected error has occurred. Could not vote";
    });        
  };

  const resetDB = () => {
    axios
    .get("http://172.234.43.240:8080/mysecretkey/resetdb")
    .then((response) => {
      console.log(response.data);
    })
    .catch((ex) => {
      const error = "An unexpected error has occurred. Could not vote";
    });        
  };

  const populateDB = () => {
    axios
    .get("http://172.234.43.240:8080/mysecretkey/populatedb")
    .then((response) => {
      console.log(response.data);
    })
    .catch((ex) => {
      const error = "An unexpected error has occurred. Could not vote";
    });        
  };

  const exportDB = () => {
    axios
    .get("http://172.234.43.240:8080/mysecretkey/export")
    .then((response) => {
      console.log(response.data);
    })
    .catch((ex) => {
      const error = "An unexpected error has occurred. Could not vote";
    });        
  };
  
  return (
    <div className="row">
    <div className="col-6">
    {ballots.map((p) => (
      <li key={p.id}>
      <p>{p.description} <button className="button" onClick={() => cueThis(p)}>Cue</button></p>            
      </li>
      ))}          
      </div>
      <div className="col-6">
        <h2>Back-end functions</h2>
        <button className="button is-green" onClick={() => resetDB()}>Delete DB (use caution)</button>
        <button className="button is-green" onClick={() => populateDB()}>Populate DB (use caution)</button>
        <button className="button is-green" onClick={() => exportDB()}>Export DB to server (use caution)</button>
        <button className="button is-green" onClick={() => toggleFade()}>Toggle Fade State</button>
      </div>
      </div>
      );
    }