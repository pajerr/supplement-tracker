import { useState, useEffect } from "react";
import axios from "axios";
import Dashboard from "./components/Dashboard";

const DisplayMagnesium = ({ magnesium }) => {
  return <div>{magnesium}</div>;
};

const DisplayDashboard = ({ dashboard }) => {
  // dashboard = FetchDashboard();
  return (
    <div>
      <ul>
        {dashboard.map((i, fakekey) => (
          <li key={fakekey}>{i.Name}</li>
        ))}
      </ul>
    </div>
  );
};

const DisplayAnecdote = ({ anecdotes, selected }) => {
  return <div>{anecdotes[selected]}</div>;
};

const Button = ({ handleClick, text }) => {
  return <button onClick={handleClick}>{text}</button>;
};

const getRandomInt = function (min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min) + min); //The maximum is exclusive and the minimum is inclusive
};

const VoteStats = ({ votes }) => {
  if (votes === 0) {
    return <div>No votes yet</div>;
  } else {
    return <div>Has {votes} votes</div>;
  }
};

const MostVotes = ({ anecdotes, points }) => {
  let mostVotesIndex = 0;
  let mostVotesAmount = 0;
  for (let i = 0; i < points.length; i++) {
    if (points[i] >= mostVotesAmount) {
      mostVotesAmount = points[i];
      mostVotesIndex = i;
    }
  }
  if (mostVotesAmount !== 0) {
    return <DisplayAnecdote anecdotes={anecdotes} selected={mostVotesIndex} />;
  } else {
    return <p>No votes yet</p>;
  }
};

const App = () => {
  const anecdotes = [
    "If it hurts, do it more often.",
    "Adding manpower to a late software project makes it later!",
    "The first 90 percent of the code accounts for the first 90 percent of the development time...The remaining 10 percent of the code accounts for the other 90 percent of the development time.",
    "Any fool can write code that a computer can understand. Good programmers write code that humans can understand.",
    "Premature optimization is the root of all evil.",
    "Debugging is twice as hard as writing the code in the first place. Therefore, if you write the code as cleverly as possible, you are, by definition, not smart enough to debug it.",
    "Programming without an extremely heavy use of console.log is same as if a doctor would refuse to use x-rays or blood tests when dianosing patients.",
  ];
  const [selected, setSelected] = useState(0);
  const [points, setPoints] = useState(new Array(6).fill(0));
  const [magnesium, setMagnesium] = useState(0);
  //const [dashboard, setDashboard] = useState([]);
  const [dashboard, setDashboard] = useState(new Array(6).fill(0));

  const handleVoteClick = () => {
    const newPoints = [...points];
    newPoints[selected] += 1;
    setPoints(newPoints);
  };

  useEffect(() => {
    /*
    console.log("effect");
    axios
      .get("http://localhost:5050/supplements/magnesium")
      .then((response) => {
        console.log("inside effect magnesium taken:", response.data, "units");
        console.log("promise fulfilled");
        setMagnesium(response.data);
      });
  }, []);
  */
    axios.get("http://localhost:5050/dashboard").then((response) => {
      console.log("inside effect dashboard:", response.data);
      console.log("promise fulfilled");
      //console.log(response.data);
      setDashboard(response.data);
      //iterate over map for debugging
      const dashboardResult = dashboard.map((i) => i.name);
      console.log("Debug print:", dashboardResult);
    });
  }, []);

  const handleTakenMagnesium = (event) => {
    event.preventDefault();
    axios
      .post("http://localhost:5050/supplements/magnesium")
      .then((response) => {
        console.log(response);
      });

    axios
      .get("http://localhost:5050/supplements/magnesium")
      .then((response) => {
        console.log(response);
        setMagnesium(response.data);
      });

    /*const newMagnesium = magnesium + 1;
    setMagnesium(newMagnesium);*/
  };

  const FetchDashboard = () => {
    axios.get("http://localhost:5050/dashboard").then((response) => {
      console.log("inside effect dashboard:", response.data);
      console.log("promise fulfilled");
      //console.log(response.data);
      setDashboard(response.data);
      //iterate over map for debugging
      const dashboardResult = dashboard.map((i) => i.name);
      console.log("Debug print:", dashboardResult);
    });
  };

  return (
    <div>
      <h1>Anecdote of the day</h1>
      <DisplayAnecdote anecdotes={anecdotes} selected={selected} />
      <VoteStats votes={points[selected]} />
      <Button
        handleClick={() => {
          setSelected(getRandomInt(0, 6));
        }}
        text="Next anecdote"
      />
      <Button handleClick={handleVoteClick} text="Vote" />
      <h1>Anecdote with the most votes</h1>
      <MostVotes anecdotes={anecdotes} points={points} />
      <h1>Magnesium taken</h1>
      <DisplayMagnesium magnesium={magnesium} />
      <Button handleClick={handleTakenMagnesium} text="Magnesium taken" />
      <h1>Dashboard</h1>
      <DisplayDashboard dashboard={dashboard} />
      <Button handleClick={FetchDashboard} text="Fetch dashboard" />
    </div>
  );
};

export default App;
