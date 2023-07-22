import { useState } from "react";

function App() {
  const [people, setPeople] = useState([]);
  return (
    <>
      <div className="text-xl">
        <form className="flex flex-row my-4 mx-2">
          <div>
            <label>Name</label>
            <input
              className="bg-lime-400 pl-4 outline-0 text-white rounded-md mx-4"
              type="text"
            />
          </div>
          <div>
            <label>DOB</label>
            <input
              className="bg-lime-400 pl-4 outline-0 text-white rounded-md mx-4"
              type="text"
            />
          </div>
          <div>
            <label>Gender</label>
            <input
              className="bg-lime-400 pl-4 outline-0 text-white rounded-md mx-4"
              type="text"
            />
          </div>
          <div>
            <label>email</label>
            <input
              className="bg-lime-400 pl-4 outline-0 text-white rounded-md mx-4"
              type="text"
            />
          </div>
        </form>
      </div>
    </>
  );
}

export default App;
