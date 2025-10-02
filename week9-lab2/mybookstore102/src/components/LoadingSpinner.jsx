import React from 'react';

const LoadingSpinner = ({ size = 'medium', text = 'กำลังโหลด...' }) => {
  const sizeClasses = {
    small: 'h-8 w-8 border-2',
    medium: 'h-12 w-12 border-3',
    large: 'h-16 w-16 border-4'
  };

  return (
    <div className="flex flex-col items-center justify-center min-h-[400px] py-12">
      <div className={`${sizeClasses[size]} border-viridian-200 border-t-viridian-600 
        rounded-full animate-spin`}></div>
      {text && (
        <p className="mt-4 text-gray-600 text-lg">{text}</p>
      )}
    </div>
  );
};

export default LoadingSpinner;