import { useState, useEffect } from "react";

import Dashboard from "./components/Dashboard";
import dashboardService from "./services/Dashboard";
import supplementService from "./services/Supplement";

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

/*
const Button = ({ handleClick, text }) => {
  return <button onClick={handleClick}>{text}</button>;
};
*/

const App = () => {
  const [newSupplement, setNewSupplement] = useState("");
  //dashboard stores information about all supplements
  const [dashboard, setDashboard] = useState(new Array().fill(0));

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
          UnitsTaken: returnedUnitsTaken,
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
    </div>
  );
};

export default App;
