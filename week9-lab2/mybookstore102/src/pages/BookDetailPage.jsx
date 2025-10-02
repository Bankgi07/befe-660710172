import React, { useState, useEffect } from 'react';
import { Link, useNavigate, useParams } from 'react-router-dom';
//import BookCard from '../components/BookCard';
//import LoadingSpinner from '../components/LoadingSpinner';
//import './BookDetailPage.css';

const BookDetailPage = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const [book, setBook] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchBook = async () => {
      try {
        setLoading(true);
        const response = await fetch(`/api/v1/books/${id}`);

        if (!response.ok) {
          throw new Error('Failed to fetch book details');
        }

        const data = await response.json();
        setBook(data);
        setError(null);
      } catch (err) {
        setError(err.message);
        console.error('Error fetching book:', err);
      } finally {
        setLoading(false);
      }
    };

    if (id) {
      fetchBook();
    }
  }, [id]);

  if (loading) {
    return (
      <div className="flex justify-center items-center min-h-screen">
        <div className="text-xl">Loading...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex flex-col justify-center items-center min-h-screen">
        <div className="text-xl text-red-600 mb-4">Error: {error}</div>
        <Link to="/books" className="text-blue-600 hover:underline">
          Go back to Book List
        </Link>
      </div>
    );
  }

  if (!book) {
    return (
      <div className="flex flex-col justify-center items-center min-h-screen">
        <div className="text-xl mb-4">Book not found</div>
        <Link to="/books" className="text-blue-600 hover:underline">
          Go back to Book List
        </Link>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <Link to="/books" className="text-blue-600 hover:underline mb-4 inline-block">
        ← Back to Book List
      </Link>

      <div className="bg-white rounded-lg shadow-lg p-6">
        <h1 className="text-3xl font-bold mb-4">{book.title}</h1>
        <div className="space-y-3">
          <p><span className="font-semibold">Author:</span> {book.author}</p>
          <p><span className="font-semibold">ISBN:</span> {book.isbn}</p>
          <p><span className="font-semibold">Year:</span> {book.year}</p>
          <p><span className="font-semibold">Price:</span> ฿{book.price}</p>
        </div>
      </div>
    </div>
  );
};

export default BookDetailPage;