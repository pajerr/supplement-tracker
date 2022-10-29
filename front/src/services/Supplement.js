import axios from "axios";
const baseUrl = "http://localhost:5050/supplements";

/*
const getAll = () => {
  const request = axios.get(baseUrl);
  return request.then((response) => response.data);
};
*/

const addTakenUnit = (supplementName) => {
  const request = axios.post(`${baseUrl}/${supplementName}`);
  return request.then((response) => response.data);
};

/*
const update = (id, newObject) => {
  const request = axios.put(`${baseUrl}/${id}`, newObject);
  return request.then((response) => response.data);
};
*/
export default { addTakenUnit };
