// Sample book data for the bookstore
export const booksData = [
  {
    id: 1,
    title: 'The Great Gatsby',
    author: 'F. Scott Fitzgerald',
    category: 'fiction',
    price: 299,
    originalPrice: 399,
    coverImage: 'images/books/gatsby.jpg',
    rating: 4.5,
    reviews: 234,
    discount: 25,
    isbn: '978-0-7432-7356-5',
    pages: 180,
    language: 'English',
    publisher: 'Scribner',
    year: 1925,
    description: 'นวนิยายคลาสสิกของอเมริกาที่เล่าเรื่องราวของเจย์ แกตส์บี้ และความฝันอเมริกันในยุค 1920s'
  },
  {
    id: 2,
    title: '1984',
    author: 'George Orwell',
    category: 'fiction',
    price: 350,
    coverImage: '/images/books/1984.jpg',
    rating: 4.8,
    reviews: 512,
    isNew: true,
    isbn: '978-0-452-28423-4',
    pages: 328,
    language: 'English',
    publisher: 'Signet Classic',
    year: 1949,
    description: 'นวนิยายดิสโทเปียที่พรรณนาถึงสังคมเผด็จการในอนาคต'
  },
  {
    id: 3,
    title: 'To Kill a Mockingbird',
    author: 'Harper Lee',
    category: 'fiction',
    price: 320,
    coverImage: '/images/books/mockingbird.jpg',
    rating: 4.6,
    reviews: 189,
    isbn: '978-0-06-112008-4',
    pages: 324,
    language: 'English',
    publisher: 'Harper Perennial',
    year: 1960,
    description: 'เรื่องราวการเติบโตและความอยุติธรรมทางเชื้อชาติในอเมริกาใต้'
  },
  {
    id: 4,
    title: 'Sapiens: A Brief History of Humankind',
    author: 'Yuval Noah Harari',
    category: 'non-fiction',
    price: 450,
    originalPrice: 550,
    coverImage: '/images/books/sapiens.jpg',
    rating: 4.7,
    reviews: 892,
    discount: 18,
    isbn: '978-0-06-231609-7',
    pages: 464,
    language: 'Thai',
    publisher: 'Harper',
    year: 2014,
    description: 'ประวัติศาสตร์ของมนุษยชาติตั้งแต่ยุคหินจนถึงปัจจุบัน'
  },
  {
    id: 5,
    title: 'The Alchemist',
    author: 'Paulo Coelho',
    category: 'fiction',
    price: 280,
    coverImage: '/images/books/alchemist.jpg',
    rating: 4.3,
    reviews: 1523,
    isbn: '978-0-06-231500-7',
    pages: 208,
    language: 'Thai',
    publisher: 'HarperOne',
    year: 1988,
    description: 'นวนิยายผจญภัยเชิงปรัชญาเกี่ยวกับการค้นหาโชคชะตาของตนเอง'
  },
  {
    id: 6,
    title: 'Thinking, Fast and Slow',
    author: 'Daniel Kahneman',
    category: 'psychology',
    price: 420,
    coverImage: '/images/books/thinking.jpg',
    rating: 4.4,
    reviews: 445,
    isNew: true,
    isbn: '978-0-374-53355-7',
    pages: 512,
    language: 'English',
    publisher: 'FSG',
    year: 2011,
    description: 'การสำรวจระบบความคิดสองระบบที่ขับเคลื่อนวิธีที่เราคิด'
  },
  {
    id: 7,
    title: 'The Art of War',
    author: 'Sun Tzu',
    category: 'history',
    price: 250,
    originalPrice: 350,
    coverImage: '/images/books/artofwar.jpg',
    rating: 4.6,
    reviews: 667,
    discount: 29,
    isbn: '978-1-59030-225-6',
    pages: 273,
    language: 'Thai',
    publisher: 'Shambhala',
    year: -500,
    description: 'ตำราพิชัยสงครามจีนโบราณที่ยังใช้ได้ในยุคปัจจุบัน'
  },
  {
    id: 8,
    title: 'Clean Code',
    author: 'Robert C. Martin',
    category: 'technology',
    price: 580,
    coverImage: '/images/books/cleancode.jpg',
    rating: 4.5,
    reviews: 234,
    isbn: '978-0-13-235088-2',
    pages: 464,
    language: 'English',
    publisher: 'Prentice Hall',
    year: 2008,
    description: 'คู่มือการเขียนโค้ดที่สะอาดและบำรุงรักษาได้'
  },
  {
    id: 9,
    title: 'The Lean Startup',
    author: 'Eric Ries',
    category: 'business',
    price: 380,
    coverImage: '/images/books/leanstartup.jpg',
    rating: 4.2,
    reviews: 556,
    isbn: '978-0-307-88789-4',
    pages: 336,
    language: 'Thai',
    publisher: 'Crown Business',
    year: 2011,
    description: 'วิธีการสร้างและบริหารสตาร์ทอัพอย่างมีประสิทธิภาพ'
  },
  {
    id: 10,
    title: 'The Power of Now',
    author: 'Eckhart Tolle',
    category: 'psychology',
    price: 320,
    coverImage: '/images/books/powerofnow.jpg',
    rating: 4.4,
    reviews: 889,
    isNew: true,
    isbn: '978-1-57731-480-6',
    pages: 236,
    language: 'Thai',
    publisher: 'New World Library',
    year: 1997,
    description: 'คู่มือการใช้ชีวิตอยู่กับปัจจุบันและการตื่นรู้ทางจิตวิญญาณ'
  },
  {
    id: 11,
    title: 'Atomic Habits',
    author: 'James Clear',
    category: 'psychology',
    price: 390,
    originalPrice: 450,
    coverImage: '/images/books/atomichabits.jpg',
    rating: 4.8,
    reviews: 2341,
    discount: 13,
    isbn: '978-0-7352-1129-2',
    pages: 320,
    language: 'Thai',
    publisher: 'Avery',
    year: 2018,
    description: 'วิธีสร้างนิสัยที่ดีและกำจัดนิสัยที่ไม่ดี'
  },
  {
    id: 12,
    title: 'The 7 Habits of Highly Effective People',
    author: 'Stephen R. Covey',
    category: 'business',
    price: 420,
    coverImage: '/images/books/7habits.jpg',
    rating: 4.5,
    reviews: 1456,
    isbn: '978-0-7432-6951-3',
    pages: 432,
    language: 'Thai',
    publisher: 'Free Press',
    year: 1989,
    description: '7 นิสัยสำหรับการพัฒนาตนเองและความสำเร็จ'
  },
  {
    id: 13,
    title: 'The Subtle Art of Not Giving a F*ck',
    author: 'Mark Manson',
    category: 'psychology',
    price: 340,
    coverImage: '/images/books/subtleart.jpg',
    rating: 4.1,
    reviews: 1789,
    isbn: '978-0-06-245771-4',
    pages: 224,
    language: 'Thai',
    publisher: 'HarperOne',
    year: 2016,
    description: 'แนวทางการใช้ชีวิตแบบตรงไปตรงมาเพื่อชีวิตที่ดีขึ้น'
  },
  {
    id: 14,
    title: 'Rich Dad Poor Dad',
    author: 'Robert T. Kiyosaki',
    category: 'business',
    price: 360,
    originalPrice: 420,
    coverImage: '/images/books/richdad.jpg',
    rating: 4.3,
    reviews: 2567,
    discount: 14,
    isbn: '978-1-61268-019-0',
    pages: 336,
    language: 'Thai',
    publisher: 'Plata Publishing',
    year: 1997,
    description: 'บทเรียนการเงินจากพ่อสองคนที่มีมุมมองต่างกัน'
  },
  {
    id: 15,
    title: 'The Da Vinci Code',
    author: 'Dan Brown',
    category: 'fiction',
    price: 380,
    coverImage: '/images/books/davinci.jpg',
    rating: 4.0,
    reviews: 3456,
    isbn: '978-0-385-50420-1',
    pages: 689,
    language: 'Thai',
    publisher: 'Doubleday',
    year: 2003,
    description: 'นวนิยายลึกลับระทึกขวัญเกี่ยวกับรหัสลับในงานศิลปะ'
  }
];

// Function to get all books
export const getAllBooks = () => {
  return booksData;
};

// Function to get a single book by ID
export const getBookById = (id) => {
  return booksData.find(book => book.id === parseInt(id));
};

// Function to get books by category
export const getBooksByCategory = (category) => {
  if (!category || category === 'all') return booksData;
  return booksData.filter(book => book.category === category);
};

// Function to search books
export const searchBooks = (query) => {
  const lowercaseQuery = query.toLowerCase();
  return booksData.filter(book => 
    book.title.toLowerCase().includes(lowercaseQuery) ||
    book.author.toLowerCase().includes(lowercaseQuery) ||
    book.category.toLowerCase().includes(lowercaseQuery)
  );
};

// Function to get featured books
export const getFeaturedBooks = (limit = 3) => {
  return booksData
    .filter(book => book.rating >= 4.5)
    .slice(0, limit);
};

// Function to get new books
export const getNewBooks = (limit = 4) => {
  return booksData
    .filter(book => book.isNew)
    .slice(0, limit);
};

// Function to get discounted books
export const getDiscountedBooks = (limit = 4) => {
  return booksData
    .filter(book => book.discount)
    .sort((a, b) => (b.discount || 0) - (a.discount || 0))
    .slice(0, limit);
};

export default booksData;