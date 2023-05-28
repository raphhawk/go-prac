import { useState, useEffect } from "react";
import { CommentCreate } from "./CommentCreate";
import { CommentList } from "./CommentList";
import axios from "axios";

export const PostList = () => {
  const [posts, setPosts] = useState({});

  const fetchPosts = async () => {
    const res = await axios.get("http://localhost:4000/posts");
    setPosts(res.data);
  };

  useEffect(() => {
    fetchPosts();
  }, []);

  const renderPosts = Object.values(posts).map((post) => {
    return (
      <div key={post.id} className="text-xl px-20 my-2 w-96 h-fit ">
        <div className="">
          <h3 className="text-2xl font-medium">{post.title}</h3>
          <CommentList postId={post.id} />
          <CommentCreate postId={post.id} />
        </div>
      </div>
    );
  });
  return <div className="grid grid-cols-2 ">{renderPosts}</div>;
};
