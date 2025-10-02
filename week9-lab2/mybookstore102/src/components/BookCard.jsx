import React, { useState } from 'react';
import { HeartIcon, ShoppingCartIcon, StarIcon } from '@heroicons/react/outline';
import { HeartIcon as HeartSolidIcon, StarIcon as StarSolidIcon } from '@heroicons/react/solid';
import { Link } from 'react-router-dom';

const BookCard = ({ book }) => {
  const [isFavorite, setIsFavorite] = useState(false);
  const [isInCart, setIsInCart] = useState(false);

  const handleAddToCart = (e) => {
    e.preventDefault();
    setIsInCart(!isInCart);
    // Add cart logic here
  };

  const handleToggleFavorite = (e) => {
    e.preventDefault();
    setIsFavorite(!isFavorite);
    // Add favorite logic here
  };

  return (
    <Link to={`/books/${book.id}`} className="block">
      <div className="bg-white rounded-xl shadow-lg overflow-hidden group 
        hover:shadow-2xl transition-all duration-300 transform hover:-translate-y-1">
        
        {/* Book Cover */}
        <div className="relative h-80 bg-gradient-to-br from-gray-100 to-gray-200">
          <img 
            src={book.coverImage || '/images/book-placeholder.jpg'} 
            alt={book.title}
            className="w-full h-full object-cover"
          />
          
          {/* Badges */}
          {book.isNew && (
            <span className="absolute top-3 left-3 bg-green-500 text-white px-3 py-1 
              rounded-full text-xs font-semibold">
              ใหม่
            </span>
          )}
          {book.discount && (
            <span className="absolute top-3 right-3 bg-red-500 text-white px-3 py-1 
              rounded-full text-xs font-semibold">
              -{book.discount}%
            </span>
          )}
          
          {/* Quick Actions - Show on Hover */}
          <div className="absolute inset-0 bg-black bg-opacity-0 group-hover:bg-opacity-40 
            transition-all duration-300 flex items-center justify-center">
            <div className="opacity-0 group-hover:opacity-100 transition-opacity duration-300 
              flex gap-3">
              <button 
                onClick={handleToggleFavorite}
                className="p-3 bg-white rounded-full hover:bg-red-50 transition-colors"
              >
                {isFavorite ? (
                  <HeartSolidIcon className="h-6 w-6 text-red-500" />
                ) : (
                  <HeartIcon className="h-6 w-6 text-gray-700" />
                )}
              </button>
              <button 
                onClick={handleAddToCart}
                className="p-3 bg-white rounded-full hover:bg-viridian-50 transition-colors"
              >
                <ShoppingCartIcon className={`h-6 w-6 ${
                  isInCart ? 'text-viridian-600' : 'text-gray-700'
                }`} />
              </button>
            </div>
          </div>
        </div>
        
        {/* Book Details */}
        <div className="p-5">
          {/* Category */}
          <p className="text-xs text-viridian-600 font-semibold uppercase tracking-wider mb-2">
            {book.category}
          </p>
          
          {/* Title */}
          <h3 className="text-lg font-bold text-gray-900 mb-1 line-clamp-1 
            group-hover:text-viridian-600 transition-colors">
            {book.title}
          </h3>
          
          {/* Author */}
          <p className="text-sm text-gray-600 mb-3">โดย {book.author}</p>
          
          {/* Rating */}
          <div className="flex items-center mb-3">
            <div className="flex text-yellow-400">
              {[...Array(5)].map((_, i) => (
                i < Math.floor(book.rating || 0) ? (
                  <StarSolidIcon key={i} className="h-4 w-4" />
                ) : (
                  <StarIcon key={i} className="h-4 w-4" />
                )
              ))}
            </div>
            <span className="text-sm text-gray-600 ml-2">
              ({book.reviews || 0} รีวิว)
            </span>
          </div>
          
          {/* Price */}
          <div className="flex items-center justify-between">
            <div>
              {book.originalPrice && book.originalPrice !== book.price && (
                <span className="text-sm text-gray-400 line-through mr-2">
                  ฿{book.originalPrice}
                </span>
              )}
              <span className="text-2xl font-bold text-viridian-600">
                ฿{book.price}
              </span>
            </div>
            
            <button 
              onClick={handleAddToCart}
              className={`px-4 py-2 rounded-lg font-semibold transition-all duration-200 
                ${isInCart 
                  ? 'bg-green-500 text-white hover:bg-green-600' 
                  : 'bg-viridian-600 text-white hover:bg-viridian-700'
                }`}>
              {isInCart ? 'ในตะกร้า' : 'เพิ่มลงตะกร้า'}
            </button>
          </div>
        </div>
      </div>
    </Link>
  );
};

export default BookCard;