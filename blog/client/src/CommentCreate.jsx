import axios from "axios";
import { useState } from "react";

export const CommentCreate = ({ postId }) => {
  const [comment, setComment] = useState("");

  const onSubmit = async (e) => {
    e.preventDefault();
    await axios.post(`http://posts.com/posts/${postId}/comments`, {
      content: comment,
    });
    setComment("");
  };

  return (
    <div>
      <form onSubmit={onSubmit}>
        <div>
          <label>Comment</label>
          <input
            value={comment}
            onChange={(e) => setComment(e.target.value)}
            className="border border-2 border-gray-500 rounded-sm focus:outline-none hover:border-blue-400 focus:bg-blue-400 focus:border-blue-400 focus: text-white px-2"
          />
        </div>
        <button className="border border-2 border-gray-500 my-2 px-2 rounded-md hover:bg-blue-400 hover:text-white hover:border-blue-400">
          Submit
        </button>
      </form>
    </div>
  );
};
