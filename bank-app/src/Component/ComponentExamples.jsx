import React from 'react';
import Welcome from './Welcome';
import Greeting from './Greeting';
import Clock from './Clock';
import UserCard from './UserCard';
import TodoList from './TodoList';

function ComponentExamples() {
    return (
        <div style={{
            maxWidth: '800px',
            margin: '0 auto',
            padding: '20px',
            fontFamily: 'Arial, sans-serif'
        }}>
            <h1 style={{
                textAlign: 'center',
                color: '#667eea',
                marginBottom: '40px'
            }}>
                ตัวอย่าง React Components
            </h1>
            
            <section style={{
                marginBottom: '40px',
                padding: '20px',
                border: '1px solid #eee',
                borderRadius: '8px',
                backgroundColor: '#fafafa'
            }}>
                <h2 style={{ color: '#333', marginBottom: '20px' }}>
                    Simple Components
                </h2>
                <Welcome />
                <Greeting />
            </section>
            
            <section style={{
                marginBottom: '40px',
                padding: '20px',
                border: '1px solid #eee',
                borderRadius: '8px',
                backgroundColor: '#fafafa'
            }}>
                <h2 style={{ color: '#333', marginBottom: '20px' }}>
                    Component with Logic
                </h2>
                <Clock />
            </section>
            
            <section style={{
                marginBottom: '40px',
                padding: '20px',
                border: '1px solid #eee',
                borderRadius: '8px',
                backgroundColor: '#fafafa'
            }}>
                <h2 style={{ color: '#333', marginBottom: '20px' }}>
                    Component Composition
                </h2>
                <UserCard />
            </section>
            
            <section style={{
                marginBottom: '40px',
                padding: '20px',
                border: '1px solid #eee',
                borderRadius: '8px',
                backgroundColor: '#fafafa'
            }}>
                <h2 style={{ color: '#333', marginBottom: '20px' }}>
                    Complex Component
                </h2>
                <TodoList />
            </section>
        </div>
    );
}

export default ComponentExamples;