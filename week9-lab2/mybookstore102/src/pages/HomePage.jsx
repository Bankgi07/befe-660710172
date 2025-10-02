import React from 'react';
import { Link } from 'react-router-dom';
import { ArrowRightIcon, BookOpenIcon, TruckIcon, ShieldCheckIcon } from '@heroicons/react/outline';

const HomePage = () => {
  const featuredBooks = [
    { id: 1, title: 'The Great Gatsby', author: 'F. Scott Fitzgerald', price: 299, image: '/images/books/gatsby.jpg' },
    { id: 2, title: '1984', author: 'George Orwell', price: 350, image: '/images/books/1984.jpg' },
    { id: 3, title: 'To Kill a Mockingbird', author: 'Harper Lee', price: 320, image: '/images/books/mockingbird.jpg' },
  ];

  const categories = [
    { name: 'นิยาย', icon: '📚', color: 'bg-blue-500' },
    { name: 'การ์ตูน', icon: '🎨', color: 'bg-pink-500' },
    { name: 'วิชาการ', icon: '🎓', color: 'bg-green-500' },
    { name: 'จิตวิทยา', icon: '🧠', color: 'bg-purple-500' },
  ];

  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="relative bg-gradient-to-r from-red-700 to-indigo-900 text-white">
        <div className="container mx-auto px-4 py-24">
          <div className="max-w-3xl">
            <h1 className="text-5xl md:text-6xl font-bold mb-6 animate-fade-in">
              ยินดีต้อนรับสู่ <span className="text-yellow-300">BookStore</span>
            </h1>
            <p className="text-xl mb-8 opacity-90">
              ค้นพบหนังสือที่คุณรัก จากคอลเลกชันมากกว่า 10,000 เล่ม
            </p>
            <div className="flex flex-col sm:flex-row gap-4">
              <Link to="/books" 
                className="inline-flex items-center justify-center px-8 py-3 bg-white 
                text-indigo-600 font-semibold rounded-lg hover:bg-gray-100 
                transform hover:scale-105 transition-all duration-200">
                เลือกซื้อหนังสือ
                <ArrowRightIcon className="ml-2 h-5 w-5" />
              </Link>
              <Link to="/categories" 
                className="inline-flex items-center justify-center px-8 py-3 
                border-2 border-white text-white font-semibold rounded-lg 
                hover:bg-white hover:text-indigo-600 transition-all duration-200">
                ดูหมวดหมู่ทั้งหมด
              </Link>
            </div>
          </div>
        </div>
        
        {/* Wave SVG */}
        <div className="absolute bottom-0 w-full">
          <svg viewBox="0 0 1440 120" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M0 120L60 110C120 100 240 80 360 70C480 60 600 60 720 65C840 70 960 80 1080 85C1200 90 1320 90 1380 90L1440 90V120H1380C1320 120 1200 120 1080 120C960 120 840 120 720 120C600 120 480 120 360 120C240 120 120 120 60 120H0V120Z" 
              fill="white"/>
          </svg>
        </div>
      </section>

      {/* Features */}
      <section className="py-16 bg-gray-50">
        <div className="container mx-auto px-4">
          <div className="grid md:grid-cols-3 gap-8">
            <div className="text-center group">
              <div className="bg-indigo-100 p-4 rounded-full w-20 h-20 mx-auto mb-4 
                group-hover:bg-indigo-200 transition-colors">
                <TruckIcon className="h-12 w-12 text-indigo-600" />
              </div>
              <h3 className="text-xl font-semibold mb-2">จัดส่งฟรี</h3>
              <p className="text-gray-600">เมื่อซื้อครบ 500 บาท</p>
            </div>
            <div className="text-center group">
              <div className="bg-green-100 p-4 rounded-full w-20 h-20 mx-auto mb-4 
                group-hover:bg-green-200 transition-colors">
                <ShieldCheckIcon className="h-12 w-12 text-green-600" />
              </div>
              <h3 className="text-xl font-semibold mb-2">รับประกันคุณภาพ</h3>
              <p className="text-gray-600">หนังสือของแท้ 100%</p>
            </div>
            <div className="text-center group">
              <div className="bg-purple-100 p-4 rounded-full w-20 h-20 mx-auto mb-4 
                group-hover:bg-purple-200 transition-colors">
                <BookOpenIcon className="h-12 w-12 text-purple-600" />
              </div>
              <h3 className="text-xl font-semibold mb-2">คอลเลกชันมากมาย</h3>
              <p className="text-gray-600">กว่า 10,000 เล่ม</p>
            </div>
          </div>
        </div>
      </section>

      {/* Categories */}
      <section className="py-16">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl font-bold text-center mb-12">หมวดหมู่ยอดนิยม</h2>
          <div className="grid grid-cols-2 md:grid-cols-4 gap-6">
            {categories.map((category, index) => (
              <Link 
                key={index}
                to={`/categories/${category.name.toLowerCase()}`}
                className="group relative overflow-hidden rounded-xl shadow-lg hover:shadow-2xl 
                  transition-all duration-300 transform hover:-translate-y-2"
              >
                <div className={`${category.color} h-40 flex flex-col items-center justify-center`}>
                  <span className="text-5xl mb-2">{category.icon}</span>
                  <h3 className="text-white font-bold text-lg">{category.name}</h3>
                </div>
                <div className="absolute inset-0 bg-black opacity-0 group-hover:opacity-20 
                  transition-opacity duration-300"></div>
              </Link>
            ))}
          </div>
        </div>
      </section>

      {/* Featured Books */}
      <section className="py-16 bg-gray-50">
        <div className="container mx-auto px-4">
          <h2 className="text-3xl font-bold text-center mb-12">หนังสือแนะนำ</h2>
          <div className="grid md:grid-cols-3 gap-8">
            {featuredBooks.map(book => (
              <div key={book.id} className="bg-white rounded-lg shadow-md hover:shadow-xl 
                transition-shadow duration-300 overflow-hidden group">
                <div className="h-64 bg-gray-200 relative overflow-hidden">
                  <img 
                    src={book.image} 
                    alt={book.title}
                    className="w-full h-full object-cover group-hover:scale-110 transition-transform duration-300"
                  />
                  <div className="absolute top-2 right-2 bg-red-500 text-white px-3 py-1 
                    rounded-full text-sm font-semibold">
                    ขายดี
                  </div>
                </div>
                <div className="p-6">
                  <h3 className="font-bold text-xl mb-2 group-hover:text-indigo-600 
                    transition-colors">{book.title}</h3>
                  <p className="text-gray-600 mb-4">โดย {book.author}</p>
                  <div className="flex justify-between items-center">
                    <span className="text-2xl font-bold text-indigo-600">฿{book.price}</span>
                    <button className="bg-indigo-600 text-white px-4 py-2 rounded-lg 
                      hover:bg-indigo-700 transition-colors">
                      เพิ่มลงตะกร้า
                    </button>
                  </div>
                </div>
              </div>
            ))}
          </div>
          <div className="text-center mt-8">
            <Link to="/books" className="inline-flex items-center text-indigo-600 
              hover:text-indigo-700 font-semibold text-lg group">
              ดูหนังสือทั้งหมด
              <ArrowRightIcon className="ml-2 h-5 w-5 group-hover:translate-x-2 
                transition-transform" />
            </Link>
          </div>
        </div>
      </section>

      {/* Newsletter */}
      <section className="py-16 bg-indigo-600">
        <div className="container mx-auto px-4 text-center">
          <h2 className="text-3xl font-bold text-white mb-4">
            รับข่าวสารและโปรโมชั่นล่าสุด
          </h2>
          <p className="text-indigo-200 mb-8">
            สมัครรับจดหมายข่าวเพื่อไม่พลาดหนังสือใหม่และส่วนลดพิเศษ
          </p>
          <form className="max-w-md mx-auto flex flex-col sm:flex-row gap-4">
            <input 
              type="email" 
              placeholder="กรอกอีเมลของคุณ"
              className="flex-1 px-6 py-3 rounded-lg focus:outline-none focus:ring-4 
                focus:ring-indigo-300"
            />
            <button type="submit" className="px-8 py-3 bg-yellow-400 text-indigo-900 
              font-semibold rounded-lg hover:bg-yellow-300 transition-colors">
              สมัครรับข่าว
            </button>
          </form>
        </div>
      </section>
    </div>
  );
};

export default HomePage;