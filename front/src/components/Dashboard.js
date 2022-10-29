const Dashboard = ({ dashboard }) => {
  return (
    <li>
      {dashboard.name} - {dashboard.dosagesTaken}
    </li>
  );
};

export default Dashboard;
