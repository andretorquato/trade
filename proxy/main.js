// Thanks Blitz
// https://medium.com/@theBliz/why-ive-built-a-coinmarketcap-proxy-e06c898b5765
const express = require("express");
const cors = require("cors");
const morgan = require("morgan");
const axios = require("axios");

require("dotenv").config();

const app = express();

app.use(morgan("tiny"));
app.use(cors());

app.get("/coinmarketcap", (req, res) => {
  let url = `https://pro-api.coinmarketcap.com/v1${req.originalUrl}`;
  axios
    .get(url, { headers: { "X-CMC_PRO_API_KEY": process.env.CMC_API_KEY } })
    .then((response) => {
      res.send(response.data);
    })
    .catch((err) => {
      console.log(err.response.data);
      res.send(err.response.data);
    });
});

app.get("/trade-api", (req, res) => {
  // let url = `http://localhost:6000${req.originalUrl}`;
  let url = `http://localhost:6000/new-market-data`;
  axios
    .get(url)
    .then((response) => {
      res.send(response.data);
    })
    .catch((err) => {
      console.log(err.response.data);
      res.send(err.response.data);
    });
    
});
const port = process.env.PORT || 5000;
app.listen(port, () => {
  console.log("Listening on port ", port);
});
