const Dashboard = ({ dashboard, handleTakenUnit }) => {
  //console.log("Dashboard debug dashboard:", dashboard);
  return (
    <tr>
      <li>Supplement: {dashboard.Name}</li>
      <li>Units taken: {dashboard.UnitsTaken}</li>
      <li>
        <button onClick={() => handleTakenUnit(dashboard.Name)}>Add</button>
      </li>
      <br></br>
    </tr>
  );
};

export default Dashboard;
