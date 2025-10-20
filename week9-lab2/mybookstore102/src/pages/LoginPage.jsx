import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { LockClosedIcon, UserIcon } from '@heroicons/react/outline';

const LoginPage = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    setError('');

    if (username === 'bookstoreadmin' && password === 'ManageBook68') {
      localStorage.setItem('isAdminAuthenticated', 'true');
      navigate('/store-manager/all-book');
    } else {
      setError('ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง');
    }
  };

  return (
    <div className="min-h-screen bg-[#1E293B] flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div>
          <div className="mx-auto h-16 w-16 bg-[#F1F5F9] rounded-full flex items-center justify-center">
            <LockClosedIcon className="h-10 w-10 text-[#16A34A]" />
          </div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-white">
            เข้าสู่ระบบ BackOffice
          </h2>
          <p className="mt-2 text-center text-sm text-gray-400">
            สำหรับผู้ดูแลระบบเท่านั้น
          </p>
        </div>

        <div className="bg-[#334155] rounded-xl shadow-2xl p-8">
          <form className="space-y-6" onSubmit={handleSubmit}>
            {error && (
              <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg">
                {error}
              </div>
            )}

            <div>
              <label htmlFor="username" className="block text-sm font-medium text-gray-300 mb-2">
                ชื่อผู้ใช้
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <UserIcon className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  id="username"
                  name="username"
                  type="text"
                  required
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  className="appearance-none block w-full pl-10 pr-3 py-3 border border-gray-600
                    rounded-lg bg-[#1E293B] text-white placeholder-gray-400 focus:outline-none focus:ring-2
                    focus:ring-green-500 focus:border-green-500"
                  placeholder="กรอกชื่อผู้ใช้"
                />
              </div>
            </div>

            <div>
              <label htmlFor="password" className="block text-sm font-medium text-gray-300 mb-2">
                รหัสผ่าน
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <LockClosedIcon className="h-5 w-5 text-gray-400" />
                </div>
                <input
                  id="password"
                  name="password"
                  type="password"
                  required
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="appearance-none block w-full pl-10 pr-3 py-3 border border-gray-600
                    rounded-lg bg-[#1E293B] text-white placeholder-gray-400 focus:outline-none focus:ring-2
                    focus:ring-green-500 focus:border-green-500"
                  placeholder="กรอกรหัสผ่าน"
                />
              </div>
            </div>

            <div>
              <button
                type="submit"
                className="w-full flex justify-center py-3 px-4 border border-transparent
                  rounded-lg shadow-sm text-sm font-medium text-white bg-green-600
                  hover:bg-green-700 focus:outline-none focus:ring-2 focus:ring-offset-2
                  focus:ring-green-500 transition-colors duration-200"
              >
                เข้าสู่ระบบ
              </button>
            </div>
          </form>
        </div>

        <div className="text-center">
          <a href="/" className="text-sm text-gray-400 hover:text-white transition-colors">
            ← กลับสู่หน้าแรก
          </a>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
