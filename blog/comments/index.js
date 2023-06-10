const express = require("express");
const bodyParser = require("body-parser");
const { randomBytes } = require("crypto");
const cors = require("cors");
const axios = require("axios");

const commentsByPostId = {};

const app = express();
app.use(bodyParser.json());
app.use(cors());

app.get("/posts/:id/comments", (req, res) => {
  res.send(commentsByPostId[req.params.id] || []);
});

app.post("/posts/:id/comments", async (req, res) => {
  const commentId = randomBytes(4).toString("hex");
  const { content } = req.body;
  const postId = req.params.id;
  const comments = commentsByPostId[postId] || [];
  const status = "pending";

  comments.push({ id: commentId, content: content, status });
  commentsByPostId[postId] = comments;

  await axios
    .post("http://event-bus-srv:4005/events", {
      type: "CommentCreated",
      data: { id: commentId, content, postId, status },
    })
    .catch((e) => console.log(e));

  res.status(201).send(comments);
});

app.post("/events", async (req, res) => {
  console.log("Event Received: ", req.body.type);

  const { type, data } = req.body;

  if (type === "CommentModerated") {
    const { id, postId, status, content } = data;
    const comments = commentsByPostId[postId];

    const comment = comments.find((comment) => comment.id === id);
    comment.status = status;

    await axios
      .post("http://event-bus-srv:4005/events", {
        type: "CommentUpdated",
        data: {
          id,
          postId,
          content,
          status,
        },
      })
      .catch((e) => console.log(e));
  }

  res.send({});
});

app.listen(4001, () => {
  console.log("Listening on port: 4001 [Comments]");
});
