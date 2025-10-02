import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';
//import BookCard from '../components/BookCard';
//import LoadingSpinner from '../components/LoadingSpinner';
//import './BookDetailPage.css';

const AboutPage = () => {
      return (
          <div>
              <h1>Welcome to the CategoryPage</h1>
              <p>This is the CategoryPage of the bookstore application.</p>
              <p>Explore our collection of books and find your next read!</p>
              <Link to="/books">Go to Book List</Link>
          </div>
      );
  }

export default AboutPage;