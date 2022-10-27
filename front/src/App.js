import { useState, useEffect } from "react";
import axios from "axios";
import Dashboard from "./components/Dashboard";
import dashboardService from "./services/Dashboard";

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

const Button = ({ handleClick, text }) => {
  return <button onClick={handleClick}>{text}</button>;
};

const App = () => {
  const [magnesium, setMagnesium] = useState(0);
  //const [dashboard, setDashboard] = useState([]);
  const [dashboard, setDashboard] = useState(new Array(6).fill(0));

  useEffect(() => {
    dashboardService.getAll().then((initialDashboard) => {
      setDashboard(initialDashboard);
    });
  }, []);
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
  /*
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
  */

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
