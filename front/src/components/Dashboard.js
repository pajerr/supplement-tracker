const Dashboard = ({ dashboard }) => {
  console.log("Dashboard debug dashboard:", dashboard);
  return (
    <li>
      {dashboard.Name} - {dashboard.DosagesTaken}{" "}
    </li>
  );
};

export default Dashboard;
