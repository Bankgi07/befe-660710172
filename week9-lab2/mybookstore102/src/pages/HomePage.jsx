import React from 'react';
import { Link } from 'react-router-dom';
import { ArrowRightIcon, BookOpenIcon, TruckIcon, ShieldCheckIcon } from '@heroicons/react/outline';
import BookCard from '../components/BookCard';
import FeaturedBooks from '../components/FeaturedBooks';
import NewBooks from '../components/NewBooks';

const HomePage  = () => {

  const categories = [
    { name: 'นิยาย', icon: '📚', color: 'bg-sky-100', slug: 'fiction' },
    { name: 'การ์ตูน', icon: '🎨', color: 'bg-rose-100', slug: 'comics' },
    { name: 'วิชาการ', icon: '🎓', color: 'bg-lime-100', slug: 'academic' },
    { name: 'จิตวิทยา', icon: '🧠', color: 'bg-indigo-100', slug: 'psychology' },
  ];

  return (
    <div className="min-h-screen">
      {/* Hero Section */}
      <section className="relative bg-gradient-to-r from-viridian-600 bg-green-700 text-white">
        <div className="container mx-auto px-4 py-24">
          <div className="max-w-3xl">
            <h1 className="text-5xl md:text-6xl font-bold mb-6 animate-fade-in">
              ยินดีต้อนรับสู่ <span className="text-yellow-300">BookStore</span>
            </h1>
            <p className="text-xl mb-8 opacity-90">
              ค้นพบหนังสือที่คุณรัก จากคอลเล็กชันมากกว่า 10,000 เล่ม
            </p>
            <div className="flex flex-col sm:flex-row gap-4">
              <Link to="/books" 
                className="inline-flex items-center justify-center px-8 py-3 bg-white 
                text-viridian-600 font-semibold rounded-lg hover:bg-gray-100 
                transform hover:scale-105 transition-all duration-200 text-green-700">
                เลือกซื้อหนังสือ
                <ArrowRightIcon className="ml-2 h-5 w-5" />
              </Link>
              <Link to="/categories" 
                className="inline-flex items-center justify-center px-8 py-3 
                border-2 border-white text-white font-semibold rounded-lg 
                hover:bg-white hover:text-viridian-600 transition-all duration-200">
                ดูหมวดหมู่ทั้งหมด
              </Link>
            </div>
          </div>
        </div>
        
        {/* Wave SVG */}
        <div className="absolute bottom-0 w-full">
          <svg viewBox="0 0 1440 120" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M0 120L60 110C120 100 240 80 360 70C480 60 600 60 720 65C840 70 960 80 1080 85C1200 90 1320 90 1380 90L1440 90V120H1380C1320 120 1200 120 1080 120C960 120 840 120 720 120C600 120 480 120 360 120C240 120 120 120 60 120H0V120Z" 
              fill="#F9FAFB"/>
          </svg>
        </div>
      </section>

      {/* Features */}
      <section className="py-16 bg-gray-50">
        <div className="container mx-auto px-4">
          <div className="grid md:grid-cols-3 gap-8">
            <div className="text-center group">
              <div className="bg-green-100 p-4 rounded-full w-20 h-20 mx-auto mb-4 
                group-hover:bg-viridian-200 transition-colors">
                <TruckIcon className="h-12 w-12 text-viridian-600" />
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
              <h3 className="text-xl font-semibold mb-2">คอลเล็กชันมากมาย</h3>
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
                to={`/categories/${category.slug}`}
                className="group relative overflow-hidden rounded-xl shadow-lg hover:shadow-2xl 
                  transition-all duration-300 transform hover:-translate-y-2"
              >
                <div className={`${category.color} h-40 flex flex-col items-center justify-center`}>
                  <span className="text-5xl mb-2">{category.icon}</span>
                  <h3 className="text-sky-950 font-bold text-lg">{category.name}</h3>
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

            <FeaturedBooks />

            <h2 className="text-3xl font-bold text-center mb-12">หนังสือใหม่</h2>
            <NewBooks />

          <div className="text-center mt-8">
            <Link to="/books" className="inline-flex items-center text-viridian-600 
              hover:text-viridian-700 font-semibold text-lg group">
              ดูหนังสือทั้งหมด
              <ArrowRightIcon className="ml-2 h-5 w-5 group-hover:translate-x-2 
                transition-transform" />
            </Link>
          </div>
        </div>
      </section>

      {/* Newsletter */}
      <section className="py-16 bg-viridian-600">
        <div className="container mx-auto px-4 text-center">
          <h2 className="text-3xl font-bold text-white mb-4">
            รับข่าวสารและโปรโมชั่นล่าสุด
          </h2>
          <p className="text-viridian-200 mb-8">
            สมัครรับจดหมายข่าวเพื่อไม่พลาดหนังสือใหม่และส่วนลดพิเศษ
          </p>
          <form className="max-w-md mx-auto flex flex-col sm:flex-row gap-4">
            <input 
              type="email" 
              placeholder="กรอกอีเมลของคุณ"
              className="flex-1 px-6 py-3 rounded-lg focus:outline-none focus:ring-4 
                focus:ring-viridian-300 text-gray-900"
            />
            <button type="submit" className="px-8 py-3 bg-yellow-400 text-viridian-900 
              font-semibold rounded-lg hover:bg-yellow-300 transition-colors">
              สมัครรับข่าว
            </button>
          </form>
        </div>
      </section>
    </div>
  );
};

export default HomePage ;