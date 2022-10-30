import { useState, useEffect } from "react";
import axios from "axios";

import Dashboard from "./components/Dashboard";
import dashboardService from "./services/Dashboard";
import supplementService from "./services/Supplement";

const DisplayMagnesium = ({ magnesium }) => {
  return <div>{magnesium}</div>;
};

const DisplayDashboard = ({ dashboard, handleTakenUnit }) => {
  return (
    //tailwind
    <div
      style={{
        textAlign: "left",
      }}
    >
      <ul>
        <table>
          <tbody>
            {dashboard.map((i, fakekey) => (
              <Dashboard
                key={fakekey}
                dashboard={i}
                handleTakenUnit={handleTakenUnit}
              />
            ))}
          </tbody>
        </table>
      </ul>
    </div>
  );
};

const Button = ({ handleClick, text }) => {
  return <button onClick={handleClick}>{text}</button>;
};

const App = () => {
  const [newSupplement, setNewSupplement] = useState("");
  //will be removed in future for testing only
  const [magnesium, setMagnesium] = useState(0);
  //dashboard stores information about all supplements
  const [dashboard, setDashboard] = useState(new Array().fill(0));

  const supplementName = "magnesium";

  //syncs changes in form field to App state
  const handleSupplementFormChange = (event) => {
    console.log(event.target.value);
    setNewSupplement(event.target.value);
  };

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

  const handleTakenUnit = (supplementName) => {
    supplementService
      .addTakenUnit(supplementName)
      .then((returnedUnitsTaken) => {
        const supplementEntry = dashboard.find(
          (i) => i.Name === supplementName
        );
        //make copy of supplement entry and update UnitsTaken with response from server
        const updatedSupplementEntry = {
          ...supplementEntry,
          DosagesTaken: returnedUnitsTaken,
        };
        console.log("updatedSupplementEntry:", updatedSupplementEntry);
        console.log("BEFORE debug dashboard:", dashboard);
        //copy old object to all entries expect updated entry that matches supplementEntry.Name of updated entry
        setDashboard(
          dashboard.map((supplementEntry) =>
            supplementEntry.Name !== supplementName
              ? supplementEntry
              : updatedSupplementEntry
          )
        );
        console.log("AFTER debug dashboard:", dashboard);
      });
  };

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
  };
  //<form onSubmit={handleTakenUnit}></form>
  return (
    <div>
      <h1>Supplemenents tracked:</h1>
      <DisplayDashboard
        dashboard={dashboard}
        handleTakenUnit={handleTakenUnit}
      />
      <h1>Add new supplement</h1>
      <form onSubmit={() => handleTakenUnit(newSupplement)}>
        <input value={newSupplement} onChange={handleSupplementFormChange} />
        <button type="submit">Save</button>
      </form>
      <h1>Magnesium taken</h1>
      <DisplayMagnesium magnesium={magnesium} />
      <button onClick={() => handleTakenUnit(supplementName)}>
        Take unit test
      </button>
      <Button
        handleClick={handleTakenMagnesium}
        text="Hard coded button Take unit"
      />
    </div>
  );
};

export default App;
