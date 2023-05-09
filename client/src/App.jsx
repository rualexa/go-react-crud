import React from "react";

function App(){
  return (
    <div>
      <button onClick={ async () => { 
        const response = await fetch('http://localhost:3000/users')
        const data = response.json()
        console.log(data)
        }}>descagar datos</button>
    </div>
  )
}

export default App