// Thanks Blitz
// https://medium.com/@theBliz/why-ive-built-a-coinmarketcap-proxy-e06c898b5765
const express = require("express");
var bodyParser = require('body-parser')
const cors = require("cors");
const morgan = require("morgan");
const axios = require("axios");

require("dotenv").config();

const app = express();
const router = express.Router();

app.use(morgan("tiny"));
app.use(bodyParser.json({limit: '50mb'}));
app.use(bodyParser.urlencoded({limit: '50mb', extended: true}));
app.use(cors());
app.use(router);

router.post("/trade-api/*", (req, res) => {
  let url = `http://localhost:6000/new-market-data`;
  axios
    .post(url, {
      data: JSON.stringify(req.body.data)
    })
    .then((response) => {
      res.send(response.data);
    })
    .catch((err) => {
      console.log(err.response.data);
      res.send(err.response.data);
    });
});
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
// app.get("/*", (req, res) => {
//   console.log(req);
//   res.send({status: "ok"})
// });

const port = process.env.PORT || 5000;
app.listen(port, () => {
  console.log("Listening on port ", port);
});
