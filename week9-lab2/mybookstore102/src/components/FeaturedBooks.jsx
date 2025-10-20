import BookCard from "./BookCard";
import { useState, useEffect } from "react";

const FeaturedBooks = () => {
  // กำหนด State สำหรับจัดการข้อมูล
  const [featuredBooks, setFeaturedBooks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        setLoading(true);
        
        // เรียก API เพื่อดึงข้อมูลหนังสือ
        const response = await fetch('/api/v1/books/');

        if (!response.ok) {
          throw new Error('Failed to fetch books');
        }

        const data = await response.json();

        // สุ่มหนังสือ 3 เล่ม
        const shuffled = [...data].sort(() => 0.5 - Math.random());
        const selected = shuffled.slice(0, 3);

        setFeaturedBooks(selected);
        setError(null);
        
      } catch (err) {
        setError(err.message);
        console.error('Error fetching books:', err);
        
      } finally {
        setLoading(false);
      }
    };

    // เรียกใช้ฟังก์ชันดึงข้อมูล
    fetchBooks();
  }, []); // [] = dependency array ว่าง = รันครั้งเดียว

  // กรณีกำลังโหลดข้อมูล
  if (loading) {
    return (
      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div className="text-center py-8 col-span-full">
          Loading...
        </div>
      </div>
    );
  }

  // กรณีเกิดข้อผิดพลาด
  if (error) {
    return (
      <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div className="text-center py-8 col-span-full text-red-600">
          Error: {error}
        </div>
      </div>
    );
  }

  // กรณีแสดงผลข้อมูลปกติ
  return (
    <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
      {featuredBooks.map(book => (
        <BookCard 
          key={book.id} 
          book={book} 
        />
      ))}
    </div>
  );
};

export default FeaturedBooks;