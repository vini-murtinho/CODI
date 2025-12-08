import React from "react";

export const Logo: React.FC = () => {
  return (
    <div className="flex items-center">
      <div className="bg-purple-700 rounded-md shadow-sm px-4 py-2 flex items-center gap-3">
        <div className="bg-white text-purple-700 font-extrabold rounded-full h-12 w-12 flex items-center justify-center text-2xl">C</div>
        <div className="text-white font-bold text-lg">vini IA</div>
      </div>
    </div>
  );
};

export default Logo;
