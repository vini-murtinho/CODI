import React from "react";

export const Logo: React.FC = () => {
  return (
    <div className="flex items-center gap-4">
      <div className="flex items-center justify-center bg-purple-700 text-white font-extrabold rounded-full h-16 w-16 text-3xl shadow-lg">C</div>
      <div className="bg-purple-700 text-white font-bold px-3 py-2 rounded-md shadow-sm">
        vini IA
      </div>
    </div>
  );
};

export default Logo;
