import { useEffect, useState } from "react";
import { useNavigate } from 'react-router-dom';

const AllBookPage = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const fetchBooks = async () => {
    try {
      setLoading(true);
      setError(null);
      const response = await fetch('/api/v1/books/');
      if (!response.ok) {
        throw new Error('Failed to fetch books');
      }
      const fetchedData = await response.json();
      setData(fetchedData);
      console.log('Books data:', fetchedData);
    } catch (err) {
      console.error('Error fetching books:', err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleNavigateAddBook = () => {
    navigate('/store-manager/add-book');
  };

  const handleNavigateEdit = (id) => {
    navigate(`/store-manager/edit-book/${id}`);
  };

  const handleDelete = async (id, title) => {
    if (window.confirm(`คุณแน่ใจหรือไม่ว่าต้องการลบหนังสือ "${title}" (ID: ${id})?`)) {
      try {
        const response = await fetch(`/api/v1/books/${id}`, {
          method: 'DELETE',
        });

        if (!response.ok) {
          throw new Error('Failed to delete book');
        }

        setData(data.filter(book => book.id !== id));
        console.log(`Book ID: ${id} deleted successfully.`);
      } catch (err) {
        console.error('Error deleting book:', err);
        alert(`เกิดข้อผิดพลาดในการลบ: ${err.message}`);
      }
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        Loading...
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center min-h-screen text-red-600">
        Error: {error}
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-900 text-white flex flex-col">
      <header className="bg-green-700 text-white p-4 shadow-md">
        <h1 className="text-2xl font-bold">จัดการหนังสือทั้งหมด</h1>
      </header>

      <main className="flex-1 container mx-auto p-6">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-3xl font-semibold">รายการหนังสือ</h2>
          <button
            onClick={handleNavigateAddBook}
            className="bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-4 rounded transition duration-200"
          >
            เพิ่มหนังสือ
          </button>
        </div>

        <div className="bg-gray-800 shadow-lg rounded-lg overflow-hidden">
          <ul>
            {data.length > 0 ? (
              data.map((book) => (
                <li
                  key={book.id}
                  className="p-4 border-b border-gray-700 last:border-b-0 flex justify-between items-center"
                >
                  <div>
                    <p className="text-sm text-gray-400">ID: {book.id}</p>
                    <p className="text-lg font-semibold">{book.title}</p>
                    <p className='text-sm text-gray-400'>โดย {book.author}</p>
                    <p className="text-sm text-gray-300">ราคา: ฿{book.price}</p>
                    <p className="text-sm text-gray-300">รายละเอียด: {book.description}</p>
                    <p className="text-sm text-gray-300">หมวดหมู่: {book.category}</p>
                    <p className="text-sm text-gray-300">ปีที่เผยแพร่: {book.year}</p>
                    <p className="text-sm text-gray-300">สถานะ: {book.status}</p>
                  </div>

                  <div className="flex space-x-2 flex-shrink-0">
                    <button
                      onClick={() => handleNavigateEdit(book.id)}
                      className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-lg transition duration-200"
                    >
                      แก้ไข
                    </button>

                    <button
                      onClick={() => handleDelete(book.id, book.title)}
                      className="bg-red-600 hover:bg-red-700 text-white font-bold py-2 px-4 rounded-lg transition duration-200"
                    >
                      ลบ
                    </button>
                  </div>
                </li>
              ))
            ) : (
              <li className="p-6 text-center text-gray-500">
                ไม่พบข้อมูลหนังสือ
              </li>
            )}
          </ul>
        </div>
      </main>

      <footer className="bg-green-700 p-4 mt-auto"></footer>
    </div>
  );
};

export default AllBookPage;
