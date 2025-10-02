import React from 'react';
import { Link } from 'react-router-dom';
import { HomeIcon, BookOpenIcon } from '@heroicons/react/outline';

const NotFound = () => {
  return (
    <div className="min-h-screen bg-gradient-to-br from-viridian-100 to-purple-100 
      flex items-center justify-center px-4">
      <div className="text-center">
        <h1 className="text-9xl font-bold text-viridian-600 mb-4">404</h1>
        <p className="text-3xl font-semibold text-gray-800 mb-4">
          ไม่พบหน้าที่คุณค้นหา
        </p>
        <p className="text-gray-600 mb-8 max-w-md mx-auto">
          อุ๊ปส์! ดูเหมือนว่าหน้าที่คุณพยายามเข้าถึงไม่มีอยู่ 
          อาจจะถูกย้ายหรือลบไปแล้ว
        </p>
        <div className="flex flex-col sm:flex-row gap-4 justify-center">
          <Link 
            to="/"
            className="inline-flex items-center px-6 py-3 bg-viridian-600 text-white 
              font-semibold rounded-lg hover:bg-viridian-700 transition-colors">
            <HomeIcon className="h-5 w-5 mr-2" />
            กลับหน้าแรก
          </Link>
          <Link 
            to="/books"
            className="inline-flex items-center px-6 py-3 bg-white text-viridian-600 
              font-semibold rounded-lg border-2 border-viridian-600 
              hover:bg-viridian-50 transition-colors">
            <BookOpenIcon className="h-5 w-5 mr-2" />
            ดูหนังสือทั้งหมด
          </Link>
        </div>
      </div>
    </div>
  );
};

export default NotFound;