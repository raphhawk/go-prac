import { PostCreate } from "./PostCreate";
import { PostList } from "./Postlist";

const App = () => {
  return (
    <div className="md:mx-40 lg:mx-80">
      <h1 className="text-3xl my-3 mx-4">Create Post</h1>
      <PostCreate />
      <hr className="h-1 bg-black bg-blue-400 my-4" />
      <h1 className="text-3xl">Posts</h1>
      <PostList />
    </div>
  );
};

export default App;
