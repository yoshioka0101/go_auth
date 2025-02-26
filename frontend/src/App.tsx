import React from "react";

const GOOGLE_AUTH_URL = "http://localhost:8080/auth/google"; // Go バックエンドの OAuth URL

const App: React.FC = () => {
  const handleLogin = () => {
    window.location.href = GOOGLE_AUTH_URL; // Google OAuth のURLへリダイレクト
  };

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h1>Google OAuth Login</h1>
      <button onClick={handleLogin} style={{ padding: "10px 20px", fontSize: "16px" }}>
        Googleでログイン
      </button>
    </div>
  );
};

export default App;
