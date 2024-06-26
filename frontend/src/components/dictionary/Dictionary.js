import React, { useEffect, useState } from 'react';
import Entries from './Entries';
import SearchKeyboard from '../keyboard/SearchKeyboard';

function Dictionary() {
    const [slp1SearchStr, setSlp1SearchStr] = useState('CAtraH');
    const [devSearchStr, setDevSearchStr] = useState('');
    const [entries, setEntries] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);
    const [entriesPerPage] = useState(3); // Adjust number per page as needed
    const totalPages = Math.ceil(entries.length / entriesPerPage);

    const handleSearch = (slp1Str, devanagariStr) => {
        setSlp1SearchStr(slp1Str);
        setDevSearchStr(devanagariStr);
    };

    const [config, setConfig] = useState({});
    useEffect(() => {
        // Fetch configuration from the environment variable
        const apiUrl = process.env.REACT_APP_API_BASE_URL;
        setConfig({ apiUrl });
    }, []);

    useEffect(() => {
        if (slp1SearchStr) {
            const fetchResults = async () => {
                const url = `${config.apiUrl}/search?slp1=${encodeURIComponent(slp1SearchStr)}&dev=${encodeURIComponent(devSearchStr)}`;
                const response = await fetch(url);
                const data = await response.json();
                setEntries(data);
                setCurrentPage(1); // Reset to first page with new data
            };

            fetchResults();
        }
    }, [slp1SearchStr, devSearchStr]);

    const nextPage = () => {
        setCurrentPage(prev => (prev < totalPages ? prev + 1 : prev));
    };

    const prevPage = () => {
        setCurrentPage(prev => (prev > 1 ? prev - 1 : prev));
    };

    return (
        <div className='entries-container'>
            <SearchKeyboard handleSearch={handleSearch} />
            <Entries
                entries={entries.slice((currentPage - 1) * entriesPerPage, currentPage * entriesPerPage)}
                devSearchStr={devSearchStr}
            />
            <div className="pagination">
                <button onClick={prevPage} disabled={currentPage === 1} className="pagination-button">←</button>
                <button onClick={nextPage} disabled={currentPage === totalPages} className="pagination-button">→</button>
            </div>
        </div>
    );
}

export default Dictionary;


