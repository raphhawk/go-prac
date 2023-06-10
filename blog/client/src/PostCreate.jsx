import React, { useState } from "react";
import axios from "axios";

export const PostCreate = () => {
  const [title, setTitle] = useState("");

  const onSubmit = async (e) => {
    e.preventDefault();
    await axios.post("http://posts.com/posts/create", {
      title: title,
    });
    setTitle("");
  };

  return (
    <div>
      <form onSubmit={onSubmit} className="px-10">
        <div className="flex flex-col">
          <label className="text-xl mb-2">Title</label>
          <input
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            className="w-60 border-2 border-gray-500 px-2 rounded-sm focus:outline-none focus:bg-blue-400 focus:border-blue-400 hover:border-blue-400 focus:text-white"
          />
        </div>
        <button className="border border-2 border-gray-500 px-4 rounded-md my-3 hover:bg-blue-400 hover:text-white hover:border-blue-400">
          Submit
        </button>
      </form>
    </div>
  );
};
