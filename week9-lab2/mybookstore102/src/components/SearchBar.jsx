import React, { useState } from 'react';
import { SearchIcon } from '@heroicons/react/outline';

const SearchBar = ({ onSearch, placeholder = "ค้นหาหนังสือ..." }) => {
  const [searchTerm, setSearchTerm] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    if (onSearch) {
      onSearch(searchTerm);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="relative">
      <div className="relative">
        <input
          type="text"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          placeholder={placeholder}
          className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg 
            focus:outline-none focus:ring-2 focus:ring-viridian-500 focus:border-transparent
            placeholder-gray-400 transition-all duration-200"
        />
        <SearchIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 
          h-5 w-5 text-gray-400" />
        <button
          type="submit"
          className="absolute right-2 top-1/2 transform -translate-y-1/2 
            px-4 py-1.5 bg-viridian-600 text-white rounded-md hover:bg-viridian-700 
            transition-colors duration-200 text-sm font-medium"
        >
          ค้นหา
        </button>
      </div>
    </form>
  );
};

export default SearchBar;