import axios from "axios";
import { useEffect, useState } from "react";

export const CommentList = ({ postId }) => {
  const [comments, setComments] = useState({});

  const fetchComments = async () => {
    const res = await axios.get(
      `http://localhost:4001/posts/${postId}/comments`
    );
    setComments(res.data);
  };

  useEffect(() => {
    fetchComments();
  }, []);

  const renderedComments = Object.values(comments).map((comment) => {
    return (
      <li key={comment.id} className="ml-5">
        {". " + comment.content}
      </li>
    );
  });
  return <ul className="text-sm">{renderedComments}</ul>;
};
