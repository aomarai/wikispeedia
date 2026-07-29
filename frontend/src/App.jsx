import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import LobbyRoom from "./components/LobbyRoom";
import GameArena from "./components/GameArena";
import LiveLeaderboard from "./components/LiveLeaderboard";

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<LobbyRoom />} />
        <Route path="/game" element={<GameArena />} />
        <Route path="/leaderboard" element={<LiveLeaderboard />} />
      </Routes>
    </BrowserRouter>
  );
};

export default App;
