import axios from "axios";
const baseUrl = "http://localhost:5050/supplements";

/*
const getAll = () => {
  const request = axios.get(baseUrl);
  return request.then((response) => response.data);
};
*/

const addTakenUnit = (supplementName) => {
  //POST works, but not rendering what we got from backend
  const request = axios.post(`${baseUrl}/${supplementName}`);
  //.post("http://localhost:5050/supplements/magnesium");
  //console.log("debug baseurl:", baseUrl);
  //console.log("deubg suppname in axios function:", supplementName);
  //console.log(request);
  /*
    .then((response) => {
      console.log(response.data);
      return response.data;
    });
    */
  return request.then((response) => response.data);
};

/*
const update = (id, newObject) => {
  const request = axios.put(`${baseUrl}/${id}`, newObject);
  return request.then((response) => response.data);
};
*/
export default { addTakenUnit };
