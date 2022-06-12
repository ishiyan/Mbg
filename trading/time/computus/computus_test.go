//nolint:testpackage
package computus

import (
	"testing"
	"time"
)

//nolint:gochecknoglobals
var (
	westernEasterSunday = []time.Time{
		date(1583, 4, 10), date(1584, 4, 1), date(1585, 4, 21), date(1586, 4, 6), date(1587, 3, 29),
		date(1588, 4, 17), date(1589, 4, 2), date(1590, 4, 22), date(1591, 4, 14), date(1592, 3, 29),
		date(1593, 4, 18), date(1594, 4, 10), date(1595, 3, 26), date(1596, 4, 14), date(1597, 4, 6),
		date(1598, 3, 22), date(1599, 4, 11), date(1600, 4, 2), date(1601, 4, 22), date(1602, 4, 7),
		date(1603, 3, 30), date(1604, 4, 18), date(1605, 4, 10), date(1606, 3, 26), date(1607, 4, 15),
		date(1608, 4, 6), date(1609, 4, 19), date(1610, 4, 11), date(1611, 4, 3), date(1612, 4, 22),
		date(1613, 4, 7), date(1614, 3, 30), date(1615, 4, 19), date(1616, 4, 3), date(1617, 3, 26),
		date(1618, 4, 15), date(1619, 3, 31), date(1620, 4, 19), date(1621, 4, 11), date(1622, 3, 27),
		date(1623, 4, 16), date(1624, 4, 7), date(1625, 3, 30), date(1626, 4, 12), date(1627, 4, 4),
		date(1628, 4, 23), date(1629, 4, 15), date(1630, 3, 31), date(1631, 4, 20), date(1632, 4, 11),
		date(1633, 3, 27), date(1634, 4, 16), date(1635, 4, 8), date(1636, 3, 23), date(1637, 4, 12),
		date(1638, 4, 4), date(1639, 4, 24), date(1640, 4, 8), date(1641, 3, 31), date(1642, 4, 20),
		date(1643, 4, 5), date(1644, 3, 27), date(1645, 4, 16), date(1646, 4, 1), date(1647, 4, 21),
		date(1648, 4, 12), date(1649, 4, 4), date(1650, 4, 17), date(1651, 4, 9), date(1652, 3, 31),
		date(1653, 4, 13), date(1654, 4, 5), date(1655, 3, 28), date(1656, 4, 16), date(1657, 4, 1),
		date(1658, 4, 21), date(1659, 4, 13), date(1660, 3, 28), date(1661, 4, 17), date(1662, 4, 9),
		date(1663, 3, 25), date(1664, 4, 13), date(1665, 4, 5), date(1666, 4, 25), date(1667, 4, 10),
		date(1668, 4, 1), date(1669, 4, 21), date(1670, 4, 6), date(1671, 3, 29), date(1672, 4, 17),
		date(1673, 4, 2), date(1674, 3, 25), date(1675, 4, 14), date(1676, 4, 5), date(1677, 4, 18),
		date(1678, 4, 10), date(1679, 4, 2), date(1680, 4, 21), date(1681, 4, 6), date(1682, 3, 29),
		date(1683, 4, 18), date(1684, 4, 2), date(1685, 4, 22), date(1686, 4, 14), date(1687, 3, 30),
		date(1688, 4, 18), date(1689, 4, 10), date(1690, 3, 26), date(1691, 4, 15), date(1692, 4, 6),
		date(1693, 3, 22), date(1694, 4, 11), date(1695, 4, 3), date(1696, 4, 22), date(1697, 4, 7),
		date(1698, 3, 30), date(1699, 4, 19), date(1700, 4, 11), date(1701, 3, 27), date(1702, 4, 16),
		date(1703, 4, 8), date(1704, 3, 23), date(1705, 4, 12), date(1706, 4, 4), date(1707, 4, 24),
		date(1708, 4, 8), date(1709, 3, 31), date(1710, 4, 20), date(1711, 4, 5), date(1712, 3, 27),
		date(1713, 4, 16), date(1714, 4, 1), date(1715, 4, 21), date(1716, 4, 12), date(1717, 3, 28),
		date(1718, 4, 17), date(1719, 4, 9), date(1720, 3, 31), date(1721, 4, 13), date(1722, 4, 5),
		date(1723, 3, 28), date(1724, 4, 16), date(1725, 4, 1), date(1726, 4, 21), date(1727, 4, 13),
		date(1728, 3, 28), date(1729, 4, 17), date(1730, 4, 9), date(1731, 3, 25), date(1732, 4, 13),
		date(1733, 4, 5), date(1734, 4, 25), date(1735, 4, 10), date(1736, 4, 1), date(1737, 4, 21),
		date(1738, 4, 6), date(1739, 3, 29), date(1740, 4, 17), date(1741, 4, 2), date(1742, 3, 25),
		date(1743, 4, 14), date(1744, 4, 5), date(1745, 4, 18), date(1746, 4, 10), date(1747, 4, 2),
		date(1748, 4, 14), date(1749, 4, 6), date(1750, 3, 29), date(1751, 4, 11), date(1752, 4, 2),
		date(1753, 4, 22), date(1754, 4, 14), date(1755, 3, 30), date(1756, 4, 18), date(1757, 4, 10),
		date(1758, 3, 26), date(1759, 4, 15), date(1760, 4, 6), date(1761, 3, 22), date(1762, 4, 11),
		date(1763, 4, 3), date(1764, 4, 22), date(1765, 4, 7), date(1766, 3, 30), date(1767, 4, 19),
		date(1768, 4, 3), date(1769, 3, 26), date(1770, 4, 15), date(1771, 3, 31), date(1772, 4, 19),
		date(1773, 4, 11), date(1774, 4, 3), date(1775, 4, 16), date(1776, 4, 7), date(1777, 3, 30),
		date(1778, 4, 19), date(1779, 4, 4), date(1780, 3, 26), date(1781, 4, 15), date(1782, 3, 31),
		date(1783, 4, 20), date(1784, 4, 11), date(1785, 3, 27), date(1786, 4, 16), date(1787, 4, 8),
		date(1788, 3, 23), date(1789, 4, 12), date(1790, 4, 4), date(1791, 4, 24), date(1792, 4, 8),
		date(1793, 3, 31), date(1794, 4, 20), date(1795, 4, 5), date(1796, 3, 27), date(1797, 4, 16),
		date(1798, 4, 8), date(1799, 3, 24), date(1800, 4, 13), date(1801, 4, 5), date(1802, 4, 18),
		date(1803, 4, 10), date(1804, 4, 1), date(1805, 4, 14), date(1806, 4, 6), date(1807, 3, 29),
		date(1808, 4, 17), date(1809, 4, 2), date(1810, 4, 22), date(1811, 4, 14), date(1812, 3, 29),
		date(1813, 4, 18), date(1814, 4, 10), date(1815, 3, 26), date(1816, 4, 14), date(1817, 4, 6),
		date(1818, 3, 22), date(1819, 4, 11), date(1820, 4, 2), date(1821, 4, 22), date(1822, 4, 7),
		date(1823, 3, 30), date(1824, 4, 18), date(1825, 4, 3), date(1826, 3, 26), date(1827, 4, 15),
		date(1828, 4, 6), date(1829, 4, 19), date(1830, 4, 11), date(1831, 4, 3), date(1832, 4, 22),
		date(1833, 4, 7), date(1834, 3, 30), date(1835, 4, 19), date(1836, 4, 3), date(1837, 3, 26),
		date(1838, 4, 15), date(1839, 3, 31), date(1840, 4, 19), date(1841, 4, 11), date(1842, 3, 27),
		date(1843, 4, 16), date(1844, 4, 7), date(1845, 3, 23), date(1846, 4, 12), date(1847, 4, 4),
		date(1848, 4, 23), date(1849, 4, 8), date(1850, 3, 31), date(1851, 4, 20), date(1852, 4, 11),
		date(1853, 3, 27), date(1854, 4, 16), date(1855, 4, 8), date(1856, 3, 23), date(1857, 4, 12),
		date(1858, 4, 4), date(1859, 4, 24), date(1860, 4, 8), date(1861, 3, 31), date(1862, 4, 20),
		date(1863, 4, 5), date(1864, 3, 27), date(1865, 4, 16), date(1866, 4, 1), date(1867, 4, 21),
		date(1868, 4, 12), date(1869, 3, 28), date(1870, 4, 17), date(1871, 4, 9), date(1872, 3, 31),
		date(1873, 4, 13), date(1874, 4, 5), date(1875, 3, 28), date(1876, 4, 16), date(1877, 4, 1),
		date(1878, 4, 21), date(1879, 4, 13), date(1880, 3, 28), date(1881, 4, 17), date(1882, 4, 9),
		date(1883, 3, 25), date(1884, 4, 13), date(1885, 4, 5), date(1886, 4, 25), date(1887, 4, 10),
		date(1888, 4, 1), date(1889, 4, 21), date(1890, 4, 6), date(1891, 3, 29), date(1892, 4, 17),
		date(1893, 4, 2), date(1894, 3, 25), date(1895, 4, 14), date(1896, 4, 5), date(1897, 4, 18),
		date(1898, 4, 10), date(1899, 4, 2), date(1900, 4, 15), date(1901, 4, 7), date(1902, 3, 30),
		date(1903, 4, 12), date(1904, 4, 3), date(1905, 4, 23), date(1906, 4, 15), date(1907, 3, 31),
		date(1908, 4, 19), date(1909, 4, 11), date(1910, 3, 27), date(1911, 4, 16), date(1912, 4, 7),
		date(1913, 3, 23), date(1914, 4, 12), date(1915, 4, 4), date(1916, 4, 23), date(1917, 4, 8),
		date(1918, 3, 31), date(1919, 4, 20), date(1920, 4, 4), date(1921, 3, 27), date(1922, 4, 16),
		date(1923, 4, 1), date(1924, 4, 20), date(1925, 4, 12), date(1926, 4, 4), date(1927, 4, 17),
		date(1928, 4, 8), date(1929, 3, 31), date(1930, 4, 20), date(1931, 4, 5), date(1932, 3, 27),
		date(1933, 4, 16), date(1934, 4, 1), date(1935, 4, 21), date(1936, 4, 12), date(1937, 3, 28),
		date(1938, 4, 17), date(1939, 4, 9), date(1940, 3, 24), date(1941, 4, 13), date(1942, 4, 5),
		date(1943, 4, 25), date(1944, 4, 9), date(1945, 4, 1), date(1946, 4, 21), date(1947, 4, 6),
		date(1948, 3, 28), date(1949, 4, 17), date(1950, 4, 9), date(1951, 3, 25), date(1952, 4, 13),
		date(1953, 4, 5), date(1954, 4, 18), date(1955, 4, 10), date(1956, 4, 1), date(1957, 4, 21),
		date(1958, 4, 6), date(1959, 3, 29), date(1960, 4, 17), date(1961, 4, 2), date(1962, 4, 22),
		date(1963, 4, 14), date(1964, 3, 29), date(1965, 4, 18), date(1966, 4, 10), date(1967, 3, 26),
		date(1968, 4, 14), date(1969, 4, 6), date(1970, 3, 29), date(1971, 4, 11), date(1972, 4, 2),
		date(1973, 4, 22), date(1974, 4, 14), date(1975, 3, 30), date(1976, 4, 18), date(1977, 4, 10),
		date(1978, 3, 26), date(1979, 4, 15), date(1980, 4, 6), date(1981, 4, 19), date(1982, 4, 11),
		date(1983, 4, 3), date(1984, 4, 22), date(1985, 4, 7), date(1986, 3, 30), date(1987, 4, 19),
		date(1988, 4, 3), date(1989, 3, 26), date(1990, 4, 15), date(1991, 3, 31), date(1992, 4, 19),
		date(1993, 4, 11), date(1994, 4, 3), date(1995, 4, 16), date(1996, 4, 7), date(1997, 3, 30),
		date(1998, 4, 12), date(1999, 4, 4), date(2000, 4, 23), date(2001, 4, 15), date(2002, 3, 31),
		date(2003, 4, 20), date(2004, 4, 11), date(2005, 3, 27), date(2006, 4, 16), date(2007, 4, 8),
		date(2008, 3, 23), date(2009, 4, 12), date(2010, 4, 4), date(2011, 4, 24), date(2012, 4, 8),
		date(2013, 3, 31), date(2014, 4, 20), date(2015, 4, 5), date(2016, 3, 27), date(2017, 4, 16),
		date(2018, 4, 1), date(2019, 4, 21), date(2020, 4, 12), date(2021, 4, 4), date(2022, 4, 17),
		date(2023, 4, 9), date(2024, 3, 31), date(2025, 4, 20), date(2026, 4, 5), date(2027, 3, 28),
		date(2028, 4, 16), date(2029, 4, 1), date(2030, 4, 21), date(2031, 4, 13), date(2032, 3, 28),
		date(2033, 4, 17), date(2034, 4, 9), date(2035, 3, 25), date(2036, 4, 13), date(2037, 4, 5),
		date(2038, 4, 25), date(2039, 4, 10), date(2040, 4, 1), date(2041, 4, 21), date(2042, 4, 6),
		date(2043, 3, 29), date(2044, 4, 17), date(2045, 4, 9), date(2046, 3, 25), date(2047, 4, 14),
		date(2048, 4, 5), date(2049, 4, 18), date(2050, 4, 10), date(2051, 4, 2), date(2052, 4, 21),
		date(2053, 4, 6), date(2054, 3, 29), date(2055, 4, 18), date(2056, 4, 2), date(2057, 4, 22),
		date(2058, 4, 14), date(2059, 3, 30), date(2060, 4, 18), date(2061, 4, 10), date(2062, 3, 26),
		date(2063, 4, 15), date(2064, 4, 6), date(2065, 3, 29), date(2066, 4, 11), date(2067, 4, 3),
		date(2068, 4, 22), date(2069, 4, 14), date(2070, 3, 30), date(2071, 4, 19), date(2072, 4, 10),
		date(2073, 3, 26), date(2074, 4, 15), date(2075, 4, 7), date(2076, 4, 19), date(2077, 4, 11),
		date(2078, 4, 3), date(2079, 4, 23), date(2080, 4, 7), date(2081, 3, 30), date(2082, 4, 19),
		date(2083, 4, 4), date(2084, 3, 26), date(2085, 4, 15), date(2086, 3, 31), date(2087, 4, 20),
		date(2088, 4, 11), date(2089, 4, 3), date(2090, 4, 16), date(2091, 4, 8), date(2092, 3, 30),
		date(2093, 4, 12), date(2094, 4, 4), date(2095, 4, 24), date(2096, 4, 15), date(2097, 3, 31),
		date(2098, 4, 20), date(2099, 4, 12), date(2100, 3, 28), date(2101, 4, 17), date(2102, 4, 9),
		date(2103, 3, 25), date(2104, 4, 13), date(2105, 4, 5), date(2106, 4, 18), date(2107, 4, 10),
		date(2108, 4, 1), date(2109, 4, 21), date(2110, 4, 6), date(2111, 3, 29), date(2112, 4, 17),
		date(2113, 4, 2), date(2114, 4, 22), date(2115, 4, 14), date(2116, 3, 29), date(2117, 4, 18),
		date(2118, 4, 10), date(2119, 3, 26), date(2120, 4, 14), date(2121, 4, 6), date(2122, 3, 29),
		date(2123, 4, 11), date(2124, 4, 2), date(2125, 4, 22), date(2126, 4, 14), date(2127, 3, 30),
		date(2128, 4, 18), date(2129, 4, 10), date(2130, 3, 26), date(2131, 4, 15), date(2132, 4, 6),
		date(2133, 4, 19), date(2134, 4, 11), date(2135, 4, 3), date(2136, 4, 22), date(2137, 4, 7),
		date(2138, 3, 30), date(2139, 4, 19), date(2140, 4, 3), date(2141, 3, 26), date(2142, 4, 15),
		date(2143, 3, 31), date(2144, 4, 19), date(2145, 4, 11), date(2146, 4, 3), date(2147, 4, 16),
		date(2148, 4, 7), date(2149, 3, 30), date(2150, 4, 12), date(2151, 4, 4), date(2152, 4, 23),
		date(2153, 4, 15), date(2154, 3, 31), date(2155, 4, 20), date(2156, 4, 11), date(2157, 3, 27),
		date(2158, 4, 16), date(2159, 4, 8), date(2160, 3, 23), date(2161, 4, 12), date(2162, 4, 4),
		date(2163, 4, 24), date(2164, 4, 8), date(2165, 3, 31), date(2166, 4, 20), date(2167, 4, 5),
		date(2168, 3, 27), date(2169, 4, 16), date(2170, 4, 1), date(2171, 4, 21), date(2172, 4, 12),
		date(2173, 4, 4), date(2174, 4, 17), date(2175, 4, 9), date(2176, 3, 31), date(2177, 4, 20),
		date(2178, 4, 5), date(2179, 3, 28), date(2180, 4, 16), date(2181, 4, 1), date(2182, 4, 21),
		date(2183, 4, 13), date(2184, 3, 28), date(2185, 4, 17), date(2186, 4, 9), date(2187, 3, 25),
		date(2188, 4, 13), date(2189, 4, 5), date(2190, 4, 25), date(2191, 4, 10), date(2192, 4, 1),
		date(2193, 4, 21), date(2194, 4, 6), date(2195, 3, 29), date(2196, 4, 17), date(2197, 4, 9),
		date(2198, 3, 25), date(2199, 4, 14), date(2200, 4, 6), date(2201, 4, 19), date(2202, 4, 11),
		date(2203, 4, 3), date(2204, 4, 22), date(2205, 4, 7), date(2206, 3, 30), date(2207, 4, 19),
		date(2208, 4, 3), date(2209, 3, 26), date(2210, 4, 15), date(2211, 3, 31), date(2212, 4, 19),
		date(2213, 4, 11), date(2214, 3, 27), date(2215, 4, 16), date(2216, 4, 7), date(2217, 3, 30),
		date(2218, 4, 12), date(2219, 4, 4), date(2220, 4, 23), date(2221, 4, 15), date(2222, 3, 31),
		date(2223, 4, 20), date(2224, 4, 11), date(2225, 3, 27), date(2226, 4, 16), date(2227, 4, 8),
		date(2228, 3, 23), date(2229, 4, 12), date(2230, 4, 4), date(2231, 4, 24), date(2232, 4, 8),
		date(2233, 3, 31), date(2234, 4, 20), date(2235, 4, 5), date(2236, 3, 27), date(2237, 4, 16),
		date(2238, 4, 1), date(2239, 4, 21), date(2240, 4, 12), date(2241, 4, 4), date(2242, 4, 17),
		date(2243, 4, 9), date(2244, 3, 31), date(2245, 4, 13), date(2246, 4, 5), date(2247, 3, 28),
		date(2248, 4, 16), date(2249, 4, 1), date(2250, 4, 21), date(2251, 4, 13), date(2252, 3, 28),
		date(2253, 4, 17), date(2254, 4, 9), date(2255, 3, 25), date(2256, 4, 13), date(2257, 4, 5),
		date(2258, 4, 25), date(2259, 4, 10), date(2260, 4, 1), date(2261, 4, 21), date(2262, 4, 6),
		date(2263, 3, 29), date(2264, 4, 17), date(2265, 4, 2), date(2266, 3, 25), date(2267, 4, 14),
		date(2268, 4, 5), date(2269, 4, 18), date(2270, 4, 10), date(2271, 4, 2), date(2272, 4, 21),
		date(2273, 4, 6), date(2274, 3, 29), date(2275, 4, 18), date(2276, 4, 2), date(2277, 4, 22),
		date(2278, 4, 14), date(2279, 3, 30), date(2280, 4, 18), date(2281, 4, 10), date(2282, 3, 26),
		date(2283, 4, 15), date(2284, 4, 6), date(2285, 3, 22), date(2286, 4, 11), date(2287, 4, 3),
		date(2288, 4, 22), date(2289, 4, 7), date(2290, 3, 30), date(2291, 4, 19), date(2292, 4, 10),
		date(2293, 3, 26), date(2294, 4, 15), date(2295, 4, 7), date(2296, 4, 19), date(2297, 4, 11),
		date(2298, 4, 3), date(2299, 4, 16), date(2300, 4, 8), date(2301, 3, 31), date(2302, 4, 20),
		date(2303, 4, 5), date(2304, 3, 27), date(2305, 4, 16), date(2306, 4, 1), date(2307, 4, 21),
		date(2308, 4, 12), date(2309, 3, 28), date(2310, 4, 17), date(2311, 4, 9), date(2312, 3, 31),
		date(2313, 4, 13), date(2314, 4, 5), date(2315, 3, 28), date(2316, 4, 16), date(2317, 4, 1),
		date(2318, 4, 21), date(2319, 4, 6), date(2320, 3, 28), date(2321, 4, 17), date(2322, 4, 9),
		date(2323, 3, 25), date(2324, 4, 13), date(2325, 4, 5), date(2326, 4, 25), date(2327, 4, 10),
		date(2328, 4, 1), date(2329, 4, 21), date(2330, 4, 6), date(2331, 3, 29), date(2332, 4, 17),
		date(2333, 4, 2), date(2334, 3, 25), date(2335, 4, 14), date(2336, 4, 5), date(2337, 4, 18),
		date(2338, 4, 10), date(2339, 3, 26), date(2340, 4, 14), date(2341, 4, 6), date(2342, 3, 29),
		date(2343, 4, 11), date(2344, 4, 2), date(2345, 4, 22), date(2346, 4, 14), date(2347, 3, 30),
		date(2348, 4, 18), date(2349, 4, 10), date(2350, 3, 26), date(2351, 4, 15), date(2352, 4, 6),
		date(2353, 3, 22), date(2354, 4, 11), date(2355, 4, 3), date(2356, 4, 22), date(2357, 4, 7),
		date(2358, 3, 30), date(2359, 4, 19), date(2360, 4, 3), date(2361, 3, 26), date(2362, 4, 15),
		date(2363, 3, 31), date(2364, 4, 19), date(2365, 4, 11), date(2366, 4, 3), date(2367, 4, 16),
		date(2368, 4, 7), date(2369, 3, 30), date(2370, 4, 19), date(2371, 4, 4), date(2372, 3, 26),
		date(2373, 4, 15), date(2374, 3, 31), date(2375, 4, 20), date(2376, 4, 11), date(2377, 3, 27),
		date(2378, 4, 16), date(2379, 4, 8), date(2380, 3, 23), date(2381, 4, 12), date(2382, 4, 4),
		date(2383, 4, 24), date(2384, 4, 8), date(2385, 3, 31), date(2386, 4, 20), date(2387, 4, 5),
		date(2388, 3, 27), date(2389, 4, 16), date(2390, 4, 8), date(2391, 3, 24), date(2392, 4, 12),
		date(2393, 4, 4), date(2394, 4, 17), date(2395, 4, 9), date(2396, 3, 31), date(2397, 4, 20),
		date(2398, 4, 5), date(2399, 3, 28), date(2400, 4, 16), date(2401, 4, 1), date(2402, 4, 21),
		date(2403, 4, 13), date(2404, 3, 28), date(2405, 4, 17), date(2406, 4, 9), date(2407, 3, 25),
		date(2408, 4, 13), date(2409, 4, 5), date(2410, 4, 25), date(2411, 4, 10), date(2412, 4, 1),
		date(2413, 4, 21), date(2414, 4, 6), date(2415, 3, 29), date(2416, 4, 17), date(2417, 4, 2),
		date(2418, 3, 25), date(2419, 4, 14), date(2420, 4, 5), date(2421, 4, 18), date(2422, 4, 10),
		date(2423, 4, 2), date(2424, 4, 21), date(2425, 4, 6), date(2426, 3, 29), date(2427, 4, 18),
		date(2428, 4, 2), date(2429, 4, 22), date(2430, 4, 14), date(2431, 3, 30), date(2432, 4, 18),
		date(2433, 4, 10), date(2434, 3, 26), date(2435, 4, 15), date(2436, 4, 6), date(2437, 3, 22),
		date(2438, 4, 11), date(2439, 4, 3), date(2440, 4, 22), date(2441, 4, 7), date(2442, 3, 30),
		date(2443, 4, 19), date(2444, 4, 10), date(2445, 3, 26), date(2446, 4, 15), date(2447, 4, 7),
		date(2448, 4, 19), date(2449, 4, 11), date(2450, 4, 3), date(2451, 4, 16), date(2452, 4, 7),
		date(2453, 3, 30), date(2454, 4, 19), date(2455, 4, 4), date(2456, 3, 26), date(2457, 4, 15),
		date(2458, 3, 31), date(2459, 4, 20), date(2460, 4, 11), date(2461, 3, 27), date(2462, 4, 16),
		date(2463, 4, 8), date(2464, 3, 30), date(2465, 4, 12), date(2466, 4, 4), date(2467, 4, 24),
		date(2468, 4, 15), date(2469, 3, 31), date(2470, 4, 20), date(2471, 4, 5), date(2472, 3, 27),
		date(2473, 4, 16), date(2474, 4, 8), date(2475, 3, 24), date(2476, 4, 12), date(2477, 4, 4),
		date(2478, 4, 24), date(2479, 4, 9), date(2480, 3, 31), date(2481, 4, 20), date(2482, 4, 5),
		date(2483, 3, 28), date(2484, 4, 16), date(2485, 4, 1), date(2486, 4, 21), date(2487, 4, 13),
		date(2488, 4, 4), date(2489, 4, 17), date(2490, 4, 9), date(2491, 3, 25), date(2492, 4, 13),
		date(2493, 4, 5), date(2494, 3, 28), date(2495, 4, 10), date(2496, 4, 1), date(2497, 4, 21),
		date(2498, 4, 13), date(2499, 3, 29), date(2500, 4, 18),
		date(4051, 4, 9), date(4052, 3, 31), date(4053, 4, 20), date(4054, 4, 12), date(4055, 3, 28),
		date(4056, 4, 16), date(4057, 4, 8), date(4058, 3, 24), date(4059, 4, 13), date(4060, 4, 4),
		date(4061, 4, 24), date(4062, 4, 9), date(4063, 4, 1), date(4064, 4, 20), date(4065, 4, 5),
		date(4066, 3, 28), date(4067, 4, 17), date(4068, 4, 1), date(4069, 4, 21), date(4070, 4, 13),
		date(4071, 3, 29), date(4072, 4, 17), date(4073, 4, 9), date(4074, 4, 1), date(4075, 4, 14),
		date(4076, 4, 5), date(4077, 3, 28), date(4078, 4, 17), date(4079, 4, 2), date(4080, 4, 21),
		date(4081, 4, 13), date(4082, 3, 29), date(4083, 4, 18), date(4084, 4, 9), date(4085, 3, 25),
		date(4086, 4, 14), date(4087, 4, 6), date(4088, 4, 25), date(4089, 4, 10), date(4090, 4, 2),
		date(4091, 4, 22), date(4092, 4, 6), date(4093, 3, 29), date(4094, 4, 18), date(4095, 4, 3),
		date(4096, 3, 25), date(4097, 4, 14), date(4098, 4, 6), date(4099, 4, 19),
	}

	orthodoxEasterSunday = []time.Time{
		date(1583, 4, 10), date(1584, 4, 29), date(1585, 4, 21), date(1586, 4, 13), date(1587, 4, 26),
		date(1588, 4, 17), date(1589, 4, 9), date(1590, 4, 29), date(1591, 4, 14), date(1592, 4, 5),
		date(1593, 4, 25), date(1594, 4, 10), date(1595, 4, 30), date(1596, 4, 21), date(1597, 4, 6),
		date(1598, 4, 26), date(1599, 4, 18), date(1600, 4, 2), date(1601, 4, 22), date(1602, 4, 14),
		date(1603, 5, 4), date(1604, 4, 18), date(1605, 4, 10), date(1606, 4, 30), date(1607, 4, 15),
		date(1608, 4, 6), date(1609, 4, 26), date(1610, 4, 18), date(1611, 4, 3), date(1612, 4, 22),
		date(1613, 4, 14), date(1614, 5, 4), date(1615, 4, 19), date(1616, 4, 10), date(1617, 4, 30),
		date(1618, 4, 15), date(1619, 4, 7), date(1620, 4, 26), date(1621, 4, 11), date(1622, 5, 1),
		date(1623, 4, 23), date(1624, 4, 7), date(1625, 4, 27), date(1626, 4, 19), date(1627, 4, 4),
		date(1628, 4, 23), date(1629, 4, 15), date(1630, 4, 7), date(1631, 4, 20), date(1632, 4, 11),
		date(1633, 5, 1), date(1634, 4, 16), date(1635, 4, 8), date(1636, 4, 27), date(1637, 4, 19),
		date(1638, 4, 4), date(1639, 4, 24), date(1640, 4, 15), date(1641, 5, 5), date(1642, 4, 20),
		date(1643, 4, 12), date(1644, 5, 1), date(1645, 4, 16), date(1646, 4, 8), date(1647, 4, 28),
		date(1648, 4, 12), date(1649, 4, 4), date(1650, 4, 24), date(1651, 4, 9), date(1652, 4, 28),
		date(1653, 4, 20), date(1654, 4, 5), date(1655, 4, 25), date(1656, 4, 16), date(1657, 4, 8),
		date(1658, 4, 21), date(1659, 4, 13), date(1660, 5, 2), date(1661, 4, 24), date(1662, 4, 9),
		date(1663, 4, 29), date(1664, 4, 20), date(1665, 4, 5), date(1666, 4, 25), date(1667, 4, 17),
		date(1668, 4, 1), date(1669, 4, 21), date(1670, 4, 13), date(1671, 5, 3), date(1672, 4, 17),
		date(1673, 4, 9), date(1674, 4, 29), date(1675, 4, 14), date(1676, 4, 5), date(1677, 4, 25),
		date(1678, 4, 10), date(1679, 4, 30), date(1680, 4, 21), date(1681, 4, 13), date(1682, 4, 26),
		date(1683, 4, 18), date(1684, 4, 9), date(1685, 4, 29), date(1686, 4, 14), date(1687, 4, 6),
		date(1688, 4, 25), date(1689, 4, 10), date(1690, 4, 30), date(1691, 4, 22), date(1692, 4, 6),
		date(1693, 4, 26), date(1694, 4, 18), date(1695, 4, 3), date(1696, 4, 22), date(1697, 4, 14),
		date(1698, 5, 4), date(1699, 4, 19), date(1700, 4, 11), date(1701, 5, 1), date(1702, 4, 16),
		date(1703, 4, 8), date(1704, 4, 27), date(1705, 4, 19), date(1706, 4, 4), date(1707, 4, 24),
		date(1708, 4, 15), date(1709, 5, 5), date(1710, 4, 20), date(1711, 4, 12), date(1712, 5, 1),
		date(1713, 4, 16), date(1714, 4, 8), date(1715, 4, 28), date(1716, 4, 12), date(1717, 5, 2),
		date(1718, 4, 24), date(1719, 4, 9), date(1720, 4, 28), date(1721, 4, 20), date(1722, 4, 5),
		date(1723, 4, 25), date(1724, 4, 16), date(1725, 4, 8), date(1726, 4, 21), date(1727, 4, 13),
		date(1728, 5, 2), date(1729, 4, 17), date(1730, 4, 9), date(1731, 4, 29), date(1732, 4, 20),
		date(1733, 4, 5), date(1734, 4, 25), date(1735, 4, 17), date(1736, 5, 6), date(1737, 4, 21),
		date(1738, 4, 13), date(1739, 5, 3), date(1740, 4, 17), date(1741, 4, 9), date(1742, 4, 29),
		date(1743, 4, 14), date(1744, 4, 5), date(1745, 4, 25), date(1746, 4, 10), date(1747, 4, 30),
		date(1748, 4, 21), date(1749, 4, 6), date(1750, 4, 26), date(1751, 4, 18), date(1752, 4, 9),
		date(1753, 4, 22), date(1754, 4, 14), date(1755, 5, 4), date(1756, 4, 25), date(1757, 4, 10),
		date(1758, 4, 30), date(1759, 4, 22), date(1760, 4, 6), date(1761, 4, 26), date(1762, 4, 18),
		date(1763, 4, 3), date(1764, 4, 22), date(1765, 4, 14), date(1766, 5, 4), date(1767, 4, 19),
		date(1768, 4, 10), date(1769, 4, 30), date(1770, 4, 15), date(1771, 4, 7), date(1772, 4, 26),
		date(1773, 4, 11), date(1774, 5, 1), date(1775, 4, 23), date(1776, 4, 14), date(1777, 4, 27),
		date(1778, 4, 19), date(1779, 4, 11), date(1780, 4, 30), date(1781, 4, 15), date(1782, 4, 7),
		date(1783, 4, 27), date(1784, 4, 11), date(1785, 5, 1), date(1786, 4, 23), date(1787, 4, 8),
		date(1788, 4, 27), date(1789, 4, 19), date(1790, 4, 4), date(1791, 4, 24), date(1792, 4, 15),
		date(1793, 5, 5), date(1794, 4, 20), date(1795, 4, 12), date(1796, 5, 1), date(1797, 4, 16),
		date(1798, 4, 8), date(1799, 4, 28), date(1800, 4, 20), date(1801, 4, 5), date(1802, 4, 25),
		date(1803, 4, 17), date(1804, 5, 6), date(1805, 4, 21), date(1806, 4, 13), date(1807, 4, 26),
		date(1808, 4, 17), date(1809, 4, 9), date(1810, 4, 29), date(1811, 4, 14), date(1812, 5, 3),
		date(1813, 4, 25), date(1814, 4, 10), date(1815, 4, 30), date(1816, 4, 21), date(1817, 4, 6),
		date(1818, 4, 26), date(1819, 4, 18), date(1820, 4, 9), date(1821, 4, 22), date(1822, 4, 14),
		date(1823, 5, 4), date(1824, 4, 18), date(1825, 4, 10), date(1826, 4, 30), date(1827, 4, 15),
		date(1828, 4, 6), date(1829, 4, 26), date(1830, 4, 18), date(1831, 5, 1), date(1832, 4, 22),
		date(1833, 4, 14), date(1834, 5, 4), date(1835, 4, 19), date(1836, 4, 10), date(1837, 4, 30),
		date(1838, 4, 15), date(1839, 4, 7), date(1840, 4, 26), date(1841, 4, 11), date(1842, 5, 1),
		date(1843, 4, 23), date(1844, 4, 7), date(1845, 4, 27), date(1846, 4, 19), date(1847, 4, 4),
		date(1848, 4, 23), date(1849, 4, 15), date(1850, 5, 5), date(1851, 4, 20), date(1852, 4, 11),
		date(1853, 5, 1), date(1854, 4, 23), date(1855, 4, 8), date(1856, 4, 27), date(1857, 4, 19),
		date(1858, 4, 4), date(1859, 4, 24), date(1860, 4, 15), date(1861, 5, 5), date(1862, 4, 20),
		date(1863, 4, 12), date(1864, 5, 1), date(1865, 4, 16), date(1866, 4, 8), date(1867, 4, 28),
		date(1868, 4, 12), date(1869, 5, 2), date(1870, 4, 24), date(1871, 4, 9), date(1872, 4, 28),
		date(1873, 4, 20), date(1874, 4, 12), date(1875, 4, 25), date(1876, 4, 16), date(1877, 4, 8),
		date(1878, 4, 28), date(1879, 4, 13), date(1880, 5, 2), date(1881, 4, 24), date(1882, 4, 9),
		date(1883, 4, 29), date(1884, 4, 20), date(1885, 4, 5), date(1886, 4, 25), date(1887, 4, 17),
		date(1888, 5, 6), date(1889, 4, 21), date(1890, 4, 13), date(1891, 5, 3), date(1892, 4, 17),
		date(1893, 4, 9), date(1894, 4, 29), date(1895, 4, 14), date(1896, 4, 5), date(1897, 4, 25),
		date(1898, 4, 17), date(1899, 4, 30), date(1900, 4, 22), date(1901, 4, 14), date(1902, 4, 27),
		date(1903, 4, 19), date(1904, 4, 10), date(1905, 4, 30), date(1906, 4, 15), date(1907, 5, 5),
		date(1908, 4, 26), date(1909, 4, 11), date(1910, 5, 1), date(1911, 4, 23), date(1912, 4, 7),
		date(1913, 4, 27), date(1914, 4, 19), date(1915, 4, 4), date(1916, 4, 23), date(1917, 4, 15),
		date(1918, 5, 5), date(1919, 4, 20), date(1920, 4, 11), date(1921, 5, 1), date(1922, 4, 16),
		date(1923, 4, 8), date(1924, 4, 27), date(1925, 4, 19), date(1926, 5, 2), date(1927, 4, 24),
		date(1928, 4, 15), date(1929, 5, 5), date(1930, 4, 20), date(1931, 4, 12), date(1932, 5, 1),
		date(1933, 4, 16), date(1934, 4, 8), date(1935, 4, 28), date(1936, 4, 12), date(1937, 5, 2),
		date(1938, 4, 24), date(1939, 4, 9), date(1940, 4, 28), date(1941, 4, 20), date(1942, 4, 5),
		date(1943, 4, 25), date(1944, 4, 16), date(1945, 5, 6), date(1946, 4, 21), date(1947, 4, 13),
		date(1948, 5, 2), date(1949, 4, 24), date(1950, 4, 9), date(1951, 4, 29), date(1952, 4, 20),
		date(1953, 4, 5), date(1954, 4, 25), date(1955, 4, 17), date(1956, 5, 6), date(1957, 4, 21),
		date(1958, 4, 13), date(1959, 5, 3), date(1960, 4, 17), date(1961, 4, 9), date(1962, 4, 29),
		date(1963, 4, 14), date(1964, 5, 3), date(1965, 4, 25), date(1966, 4, 10), date(1967, 4, 30),
		date(1968, 4, 21), date(1969, 4, 13), date(1970, 4, 26), date(1971, 4, 18), date(1972, 4, 9),
		date(1973, 4, 29), date(1974, 4, 14), date(1975, 5, 4), date(1976, 4, 25), date(1977, 4, 10),
		date(1978, 4, 30), date(1979, 4, 22), date(1980, 4, 6), date(1981, 4, 26), date(1982, 4, 18),
		date(1983, 5, 8), date(1984, 4, 22), date(1985, 4, 14), date(1986, 5, 4), date(1987, 4, 19),
		date(1988, 4, 10), date(1989, 4, 30), date(1990, 4, 15), date(1991, 4, 7), date(1992, 4, 26),
		date(1993, 4, 18), date(1994, 5, 1), date(1995, 4, 23), date(1996, 4, 14), date(1997, 4, 27),
		date(1998, 4, 19), date(1999, 4, 11), date(2000, 4, 30), date(2001, 4, 15), date(2002, 5, 5),
		date(2003, 4, 27), date(2004, 4, 11), date(2005, 5, 1), date(2006, 4, 23), date(2007, 4, 8),
		date(2008, 4, 27), date(2009, 4, 19), date(2010, 4, 4), date(2011, 4, 24), date(2012, 4, 15),
		date(2013, 5, 5), date(2014, 4, 20), date(2015, 4, 12), date(2016, 5, 1), date(2017, 4, 16),
		date(2018, 4, 8), date(2019, 4, 28), date(2020, 4, 19), date(2021, 5, 2), date(2022, 4, 24),
		date(2023, 4, 16), date(2024, 5, 5), date(2025, 4, 20), date(2026, 4, 12), date(2027, 5, 2),
		date(2028, 4, 16), date(2029, 4, 8), date(2030, 4, 28), date(2031, 4, 13), date(2032, 5, 2),
		date(2033, 4, 24), date(2034, 4, 9), date(2035, 4, 29), date(2036, 4, 20), date(2037, 4, 5),
		date(2038, 4, 25), date(2039, 4, 17), date(2040, 5, 6), date(2041, 4, 21), date(2042, 4, 13),
		date(2043, 5, 3), date(2044, 4, 24), date(2045, 4, 9), date(2046, 4, 29), date(2047, 4, 21),
		date(2048, 4, 5), date(2049, 4, 25), date(2050, 4, 17), date(2051, 5, 7), date(2052, 4, 21),
		date(2053, 4, 13), date(2054, 5, 3), date(2055, 4, 18), date(2056, 4, 9), date(2057, 4, 29),
		date(2058, 4, 14), date(2059, 5, 4), date(2060, 4, 25), date(2061, 4, 10), date(2062, 4, 30),
		date(2063, 4, 22), date(2064, 4, 13), date(2065, 4, 26), date(2066, 4, 18), date(2067, 4, 10),
		date(2068, 4, 29), date(2069, 4, 14), date(2070, 5, 4), date(2071, 4, 19), date(2072, 4, 10),
		date(2073, 4, 30), date(2074, 4, 22), date(2075, 4, 7), date(2076, 4, 26), date(2077, 4, 18),
		date(2078, 5, 8), date(2079, 4, 23), date(2080, 4, 14), date(2081, 5, 4), date(2082, 4, 19),
		date(2083, 4, 11), date(2084, 4, 30), date(2085, 4, 15), date(2086, 4, 7), date(2087, 4, 27),
		date(2088, 4, 18), date(2089, 5, 1), date(2090, 4, 23), date(2091, 4, 8), date(2092, 4, 27),
		date(2093, 4, 19), date(2094, 4, 11), date(2095, 4, 24), date(2096, 4, 15), date(2097, 5, 5),
		date(2098, 4, 27), date(2099, 4, 12), date(2100, 5, 2), date(2101, 4, 24), date(2102, 4, 9),
		date(2103, 4, 29), date(2104, 4, 20), date(2105, 4, 5), date(2106, 4, 25), date(2107, 4, 17),
		date(2108, 5, 6), date(2109, 4, 21), date(2110, 4, 13), date(2111, 5, 3), date(2112, 4, 17),
		date(2113, 4, 9), date(2114, 4, 29), date(2115, 4, 14), date(2116, 5, 3), date(2117, 4, 25),
		date(2118, 4, 17), date(2119, 4, 30), date(2120, 4, 21), date(2121, 4, 13), date(2122, 5, 3),
		date(2123, 4, 18), date(2124, 4, 9), date(2125, 4, 29), date(2126, 4, 14), date(2127, 5, 4),
		date(2128, 4, 25), date(2129, 4, 10), date(2130, 4, 30), date(2131, 4, 22), date(2132, 4, 6),
		date(2133, 4, 26), date(2134, 4, 18), date(2135, 5, 8), date(2136, 4, 22), date(2137, 4, 14),
		date(2138, 5, 4), date(2139, 4, 19), date(2140, 4, 10), date(2141, 4, 30), date(2142, 4, 22),
		date(2143, 4, 7), date(2144, 4, 26), date(2145, 4, 18), date(2146, 5, 8), date(2147, 4, 23),
		date(2148, 4, 14), date(2149, 5, 4), date(2150, 4, 19), date(2151, 4, 11), date(2152, 4, 30),
		date(2153, 4, 15), date(2154, 5, 5), date(2155, 4, 27), date(2156, 4, 11), date(2157, 5, 1),
		date(2158, 4, 23), date(2159, 4, 8), date(2160, 4, 27), date(2161, 4, 19), date(2162, 4, 11),
		date(2163, 4, 24), date(2164, 4, 15), date(2165, 5, 5), date(2166, 4, 20), date(2167, 4, 12),
		date(2168, 5, 1), date(2169, 4, 23), date(2170, 4, 8), date(2171, 4, 28), date(2172, 4, 19),
		date(2173, 5, 9), date(2174, 4, 24), date(2175, 4, 16), date(2176, 5, 5), date(2177, 4, 20),
		date(2178, 4, 12), date(2179, 5, 2), date(2180, 4, 16), date(2181, 4, 8), date(2182, 4, 28),
		date(2183, 4, 13), date(2184, 5, 2), date(2185, 4, 24), date(2186, 4, 9), date(2187, 4, 29),
		date(2188, 4, 20), date(2189, 4, 12), date(2190, 4, 25), date(2191, 4, 17), date(2192, 5, 6),
		date(2193, 4, 28), date(2194, 4, 13), date(2195, 5, 3), date(2196, 4, 24), date(2197, 4, 9),
		date(2198, 4, 29), date(2199, 4, 21), date(2200, 4, 6), date(2201, 4, 26), date(2202, 4, 18),
		date(2203, 5, 8), date(2204, 4, 22), date(2205, 4, 14), date(2206, 5, 4), date(2207, 4, 19),
		date(2208, 4, 10), date(2209, 4, 30), date(2210, 4, 15), date(2211, 5, 5), date(2212, 4, 26),
		date(2213, 4, 18), date(2214, 5, 1), date(2215, 4, 23), date(2216, 4, 14), date(2217, 5, 4),
		date(2218, 4, 19), date(2219, 4, 11), date(2220, 4, 30), date(2221, 4, 15), date(2222, 5, 5),
		date(2223, 4, 27), date(2224, 4, 11), date(2225, 5, 1), date(2226, 4, 23), date(2227, 4, 8),
		date(2228, 4, 27), date(2229, 4, 19), date(2230, 5, 9), date(2231, 4, 24), date(2232, 4, 15),
		date(2233, 5, 5), date(2234, 4, 20), date(2235, 4, 12), date(2236, 5, 1), date(2237, 4, 23),
		date(2238, 4, 8), date(2239, 4, 28), date(2240, 4, 19), date(2241, 5, 9), date(2242, 4, 24),
		date(2243, 4, 16), date(2244, 5, 5), date(2245, 4, 20), date(2246, 4, 12), date(2247, 5, 2),
		date(2248, 4, 16), date(2249, 5, 6), date(2250, 4, 28), date(2251, 4, 13), date(2252, 5, 2),
		date(2253, 4, 24), date(2254, 4, 9), date(2255, 4, 29), date(2256, 4, 20), date(2257, 4, 12),
		date(2258, 4, 25), date(2259, 4, 17), date(2260, 5, 6), date(2261, 4, 21), date(2262, 4, 13),
		date(2263, 5, 3), date(2264, 4, 24), date(2265, 4, 9), date(2266, 4, 29), date(2267, 4, 21),
		date(2268, 5, 10), date(2269, 4, 25), date(2270, 4, 17), date(2271, 5, 7), date(2272, 4, 21),
		date(2273, 4, 13), date(2274, 5, 3), date(2275, 4, 18), date(2276, 4, 9), date(2277, 4, 29),
		date(2278, 4, 14), date(2279, 5, 4), date(2280, 4, 25), date(2281, 4, 10), date(2282, 4, 30),
		date(2283, 4, 22), date(2284, 4, 13), date(2285, 4, 26), date(2286, 4, 18), date(2287, 5, 8),
		date(2288, 4, 29), date(2289, 4, 14), date(2290, 5, 4), date(2291, 4, 26), date(2292, 4, 10),
		date(2293, 4, 30), date(2294, 4, 22), date(2295, 4, 7), date(2296, 4, 26), date(2297, 4, 18),
		date(2298, 5, 8), date(2299, 4, 23), date(2300, 4, 15), date(2301, 5, 5), date(2302, 4, 20),
		date(2303, 4, 12), date(2304, 5, 1), date(2305, 4, 16), date(2306, 5, 6), date(2307, 4, 28),
		date(2308, 4, 19), date(2309, 5, 2), date(2310, 4, 24), date(2311, 4, 16), date(2312, 5, 5),
		date(2313, 4, 20), date(2314, 4, 12), date(2315, 5, 2), date(2316, 4, 16), date(2317, 5, 6),
		date(2318, 4, 28), date(2319, 4, 13), date(2320, 5, 2), date(2321, 4, 24), date(2322, 4, 9),
		date(2323, 4, 29), date(2324, 4, 20), date(2325, 5, 10), date(2326, 4, 25), date(2327, 4, 17),
		date(2328, 5, 6), date(2329, 4, 21), date(2330, 4, 13), date(2331, 5, 3), date(2332, 4, 24),
		date(2333, 4, 9), date(2334, 4, 29), date(2335, 4, 21), date(2336, 5, 10), date(2337, 4, 25),
		date(2338, 4, 17), date(2339, 4, 30), date(2340, 4, 21), date(2341, 4, 13), date(2342, 5, 3),
		date(2343, 4, 18), date(2344, 5, 7), date(2345, 4, 29), date(2346, 4, 14), date(2347, 5, 4),
		date(2348, 4, 25), date(2349, 4, 10), date(2350, 4, 30), date(2351, 4, 22), date(2352, 4, 13),
		date(2353, 4, 26), date(2354, 4, 18), date(2355, 5, 8), date(2356, 4, 22), date(2357, 4, 14),
		date(2358, 5, 4), date(2359, 4, 19), date(2360, 4, 10), date(2361, 4, 30), date(2362, 4, 22),
		date(2363, 5, 5), date(2364, 4, 26), date(2365, 4, 18), date(2366, 5, 8), date(2367, 4, 23),
		date(2368, 4, 14), date(2369, 5, 4), date(2370, 4, 19), date(2371, 4, 11), date(2372, 4, 30),
		date(2373, 4, 15), date(2374, 5, 5), date(2375, 4, 27), date(2376, 4, 11), date(2377, 5, 1),
		date(2378, 4, 23), date(2379, 4, 8), date(2380, 4, 27), date(2381, 4, 19), date(2382, 5, 9),
		date(2383, 4, 24), date(2384, 4, 15), date(2385, 5, 5), date(2386, 4, 27), date(2387, 4, 12),
		date(2388, 5, 1), date(2389, 4, 23), date(2390, 4, 8), date(2391, 4, 28), date(2392, 4, 19),
		date(2393, 5, 9), date(2394, 4, 24), date(2395, 4, 16), date(2396, 5, 5), date(2397, 4, 20),
		date(2398, 4, 12), date(2399, 5, 2), date(2400, 4, 16), date(2401, 5, 6), date(2402, 4, 28),
		date(2403, 4, 13), date(2404, 5, 2), date(2405, 4, 24), date(2406, 4, 16), date(2407, 4, 29),
		date(2408, 4, 20), date(2409, 4, 12), date(2410, 5, 2), date(2411, 4, 17), date(2412, 5, 6),
		date(2413, 4, 28), date(2414, 4, 13), date(2415, 5, 3), date(2416, 4, 24), date(2417, 4, 9),
		date(2418, 4, 29), date(2419, 4, 21), date(2420, 5, 10), date(2421, 4, 25), date(2422, 4, 17),
		date(2423, 5, 7), date(2424, 4, 21), date(2425, 4, 13), date(2426, 5, 3), date(2427, 4, 18),
		date(2428, 4, 9), date(2429, 4, 29), date(2430, 4, 21), date(2431, 5, 4), date(2432, 4, 25),
		date(2433, 4, 17), date(2434, 4, 30), date(2435, 4, 22), date(2436, 4, 13), date(2437, 5, 3),
		date(2438, 4, 18), date(2439, 5, 8), date(2440, 4, 29), date(2441, 4, 14), date(2442, 5, 4),
		date(2443, 4, 26), date(2444, 4, 10), date(2445, 4, 30), date(2446, 4, 22), date(2447, 4, 7),
		date(2448, 4, 26), date(2449, 4, 18), date(2450, 5, 8), date(2451, 4, 23), date(2452, 4, 14),
		date(2453, 5, 4), date(2454, 4, 19), date(2455, 4, 11), date(2456, 4, 30), date(2457, 4, 22),
		date(2458, 5, 5), date(2459, 4, 27), date(2460, 4, 18), date(2461, 5, 8), date(2462, 4, 23),
		date(2463, 4, 15), date(2464, 5, 4), date(2465, 4, 19), date(2466, 4, 11), date(2467, 5, 1),
		date(2468, 4, 15), date(2469, 5, 5), date(2470, 4, 27), date(2471, 4, 12), date(2472, 5, 1),
		date(2473, 4, 23), date(2474, 4, 8), date(2475, 4, 28), date(2476, 4, 19), date(2477, 5, 9),
		date(2478, 4, 24), date(2479, 4, 16), date(2480, 5, 5), date(2481, 4, 27), date(2482, 4, 12),
		date(2483, 5, 2), date(2484, 4, 23), date(2485, 4, 8), date(2486, 4, 28), date(2487, 4, 20),
		date(2488, 5, 9), date(2489, 4, 24), date(2490, 4, 16), date(2491, 5, 6), date(2492, 4, 20),
		date(2493, 4, 12), date(2494, 5, 2), date(2495, 4, 17), date(2496, 5, 6), date(2497, 4, 28),
		date(2498, 4, 13), date(2499, 5, 3), date(2500, 4, 25),
	}
)

func TestEasterSunday(t *testing.T) {
	const name = "EasterSunday"

	t.Parallel()
	verifyRange(t, westernEasterSunday, EasterSunday, name)
	verifyYearRange(t, easterSundayPerYearFirstYear, easterSundayPerYearLastYear+1000, EasterSunday, name)
	verifyYearError(t, EasterSunday, easterSundayPerYearFirstYear-1, name)
}

func TestOrthodoxEasterSunday(t *testing.T) {
	const name = "OrthodoxEasterSunday"

	t.Parallel()
	verifyRange(t, orthodoxEasterSunday, OrthodoxEasterSunday, name)
	verifyYearRange(t, easterSundayPerYearFirstYear, easterSundayPerYearLastYear, OrthodoxEasterSunday, name)
	verifyYearError(t, OrthodoxEasterSunday, easterSundayPerYearFirstYear-1, name)
	verifyYearError(t, OrthodoxEasterSunday, easterSundayPerYearLastYear+1, name)
}

func TestEasterSundayYearDay(t *testing.T) {
	const name = "EasterSundayYearDay"

	t.Parallel()
	verifyYearDayRange(t, westernEasterSunday, EasterSundayYearDay, name)
	verifyYearDayError(t, EasterSundayYearDay, easterSundayPerYearFirstYear-1, name)
}

func TestOrthodoxEasterSundayYearDay(t *testing.T) {
	const name = "OrthodoxEasterSundayYearDay"

	t.Parallel()
	verifyYearDayRange(t, orthodoxEasterSunday, OrthodoxEasterSundayYearDay, name)
	verifyYearDayError(t, OrthodoxEasterSundayYearDay, easterSundayPerYearFirstYear-1, name)
	verifyYearDayError(t, OrthodoxEasterSundayYearDay, easterSundayPerYearLastYear+1, name)
}

func TestIsEasterSunday(t *testing.T) {
	const name = "EasterSunday"

	t.Parallel()
	verifyRangeIsHoliday(t, westernEasterSunday, 0, IsEasterSunday, name)
	verifyRangeIsNotHoliday(t, westernEasterSunday, 0, IsEasterSunday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsEasterSunday, name)
}

func TestIsOrthodoxEasterSunday(t *testing.T) {
	const name = "OrthodoxEasterSunday"

	t.Parallel()
	verifyRangeIsHoliday(t, orthodoxEasterSunday, 0, IsOrthodoxEasterSunday, name)
	verifyRangeIsNotHoliday(t, orthodoxEasterSunday, 0, IsOrthodoxEasterSunday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsOrthodoxEasterSunday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearLastYear+1, IsOrthodoxEasterSunday, name)
}

func TestIsEasterMonday(t *testing.T) {
	const name = "EasterMonday"

	t.Parallel()
	verifyRangeIsHoliday(t, westernEasterSunday, 1, IsEasterMonday, name)
	verifyRangeIsNotHoliday(t, westernEasterSunday, 1, IsEasterMonday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsEasterMonday, name)
}

func TestIsOrthodoxEasterMonday(t *testing.T) {
	const name = "OrthodoxEasterSunday"

	t.Parallel()
	verifyRangeIsHoliday(t, orthodoxEasterSunday, 1, IsOrthodoxEasterMonday, name)
	verifyRangeIsNotHoliday(t, orthodoxEasterSunday, 1, IsOrthodoxEasterMonday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsOrthodoxEasterMonday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearLastYear+1, IsOrthodoxEasterMonday, name)
}

func TestIsShroveTuesday(t *testing.T) {
	const name = "ShroveTuesday"

	// Taken from https://en.wikipedia.org/wiki/Shrove_Tuesday.
	tests := []time.Time{
		date(2016, 2, 9), date(2017, 2, 28), date(2018, 2, 13), date(2019, 3, 5), date(2020, 2, 25),
		date(2021, 2, 16), date(2022, 3, 1), date(2023, 2, 21), date(2024, 2, 13), date(2025, 3, 4),
		date(2026, 2, 17), date(2027, 2, 9), date(2028, 2, 29), date(2029, 2, 13), date(2030, 3, 5),
		date(2031, 2, 25), date(2032, 2, 10), date(2033, 3, 1), date(2034, 2, 21), date(2035, 2, 6),
		date(2036, 2, 26), date(2037, 2, 17), date(2038, 3, 9), date(2039, 2, 22), date(2040, 2, 14),
		date(2041, 3, 5), date(2042, 2, 18), date(2043, 2, 10), date(2044, 3, 1), date(2045, 2, 21),
		date(2046, 2, 6), date(2047, 2, 26), date(2048, 2, 18), date(2049, 3, 2), date(2050, 2, 22),
		date(2061, 2, 22), date(2062, 2, 7), date(2063, 2, 27), date(2064, 2, 19), date(2065, 2, 10),
		date(2066, 2, 23), date(2067, 2, 15), date(2068, 3, 6), date(2069, 2, 26), date(2070, 2, 11),
		date(2071, 3, 3), date(2072, 2, 23), date(2073, 2, 7), date(2074, 2, 27), date(2075, 2, 19),
		date(2076, 3, 3), date(2077, 2, 23), date(2078, 2, 15), date(2079, 3, 7), date(2080, 2, 20),
		date(2081, 2, 11), date(2082, 3, 3), date(2083, 2, 16), date(2084, 2, 8), date(2085, 2, 27),
		date(2086, 2, 12), date(2087, 3, 4), date(2088, 2, 24), date(2089, 2, 15), date(2090, 2, 28),
		date(2091, 2, 20), date(2092, 2, 12), date(2093, 2, 24), date(2094, 2, 16), date(2095, 3, 8),
		date(2096, 2, 28), date(2097, 2, 12), date(2098, 3, 4), date(2099, 2, 24), date(2100, 2, 9),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsShroveTuesday, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsShroveTuesday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsShroveTuesday, name)
}

func TestIsAshWednesday(t *testing.T) {
	const name = "AshWednesday"

	// Taken from https://en.wikipedia.org/wiki/Ash_Wednesday.
	tests := []time.Time{
		date(2016, 2, 10), date(2017, 3, 1), date(2018, 2, 14), date(2019, 3, 6), date(2020, 2, 26),
		date(2021, 2, 17), date(2022, 3, 2), date(2023, 2, 22), date(2024, 2, 14), date(2025, 3, 5),
		date(2026, 2, 18), date(2027, 2, 10), date(2028, 3, 1), date(2029, 2, 14), date(2030, 3, 6),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsAshWednesday, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsAshWednesday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsAshWednesday, name)
}

func TestIsPalmSunday(t *testing.T) {
	const name = "PalmSunday"

	// Taken from https://en.wikipedia.org/wiki/Palm_Sunday.
	tests := []time.Time{
		date(2009, 4, 5), date(2010, 3, 28), date(2011, 4, 17), date(2012, 4, 1), date(2013, 3, 24),
		date(2014, 4, 13), date(2015, 3, 29), date(2016, 3, 20), date(2017, 4, 9), date(2018, 3, 25),
		date(2019, 4, 14), date(2020, 4, 5), date(2021, 3, 28), date(2022, 4, 10), date(2023, 4, 2),
		date(2024, 3, 24), date(2025, 4, 13), date(2026, 3, 29), date(2027, 3, 21), date(2028, 4, 9),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsPalmSunday, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsPalmSunday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsPalmSunday, name)
}

func TestIsMaundyThursday(t *testing.T) {
	const name = "MaundyThursday"

	// Taken from https://en.wikipedia.org/wiki/Maundy_Thursday.
	tests := []time.Time{
		date(2015, 4, 2), date(2016, 3, 24), date(2017, 4, 13), date(2018, 3, 29), date(2019, 4, 18),
		date(2020, 4, 9), date(2021, 4, 1), date(2022, 4, 14), date(2023, 4, 6), date(2024, 3, 28),
		date(2025, 4, 17), date(2026, 4, 2), date(2027, 3, 25), date(2028, 4, 13), date(2029, 3, 29),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsMaundyThursday, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsMaundyThursday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsMaundyThursday, name)
}

func TestIsGoodFriday(t *testing.T) {
	const name = "GoodFriday"

	// Taken from https://en.wikipedia.org/wiki/Good_Friday.
	tests := []time.Time{
		date(2015, 4, 3), date(2016, 3, 25), date(2017, 4, 14), date(2018, 3, 30), date(2019, 4, 19),
		date(2020, 4, 10), date(2021, 4, 2), date(2022, 4, 15), date(2023, 4, 7), date(2024, 3, 29),
		date(2025, 4, 18), date(2026, 4, 3), date(2027, 3, 26), date(2028, 4, 14), date(2029, 3, 30),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsGoodFriday, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsGoodFriday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsGoodFriday, name)
}

func TestIsAscensionThursday(t *testing.T) {
	const name = "AscensionThursday"

	// Taken from https://en.wikipedia.org/wiki/Feast_of_the_Ascension.
	tests := []time.Time{
		date(2000, 6, 1), date(2001, 5, 24), date(2002, 5, 9), date(2003, 5, 29), date(2004, 5, 20),
		date(2005, 5, 5), date(2006, 5, 25), date(2007, 5, 17), date(2008, 5, 1), date(2009, 5, 21),
		date(2010, 5, 13), date(2011, 6, 2), date(2012, 5, 17), date(2013, 5, 9), date(2014, 5, 29),
		date(2015, 5, 14), date(2016, 5, 5), date(2017, 5, 25), date(2018, 5, 10), date(2019, 5, 30),
		date(2020, 5, 21), date(2021, 5, 13), date(2022, 5, 26), date(2023, 5, 18), date(2024, 5, 9),
		date(2025, 5, 29), date(2026, 5, 14), date(2027, 5, 6), date(2028, 5, 25), date(2029, 5, 10),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsAscensionThursday, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsAscensionThursday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsAscensionThursday, name)
}

func TestIsWhitSunday(t *testing.T) {
	const name = "WhitSunday"

	// Taken from https://en.wikipedia.org/wiki/Pentecost.
	tests := []time.Time{
		date(2015, 5, 24), date(2016, 5, 15), date(2017, 6, 4), date(2018, 5, 20), date(2019, 6, 9),
		date(2020, 5, 31), date(2021, 5, 23), date(2022, 6, 5), date(2023, 5, 28), date(2024, 5, 19),
		date(2025, 6, 8), date(2026, 5, 24), date(2027, 5, 16), date(2028, 6, 4), date(2029, 5, 20),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsWhitSunday, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsWhitSunday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsWhitSunday, name)
}

func TestIsWhitMonday(t *testing.T) {
	const name = "WhitMonday"

	// Taken from https://en.wikipedia.org/wiki/Whit_Monday.
	tests := []time.Time{
		date(2000, 6, 12), date(2001, 6, 4), date(2002, 5, 20), date(2003, 6, 9), date(2004, 5, 31),
		date(2005, 5, 16), date(2006, 6, 5), date(2007, 5, 28), date(2008, 5, 12), date(2009, 6, 1),
		date(2010, 5, 24), date(2011, 6, 13), date(2012, 5, 28), date(2013, 5, 20), date(2014, 6, 9),
		date(2015, 5, 25), date(2016, 5, 16), date(2017, 6, 5), date(2018, 5, 21), date(2019, 6, 10),
		date(2020, 6, 1), date(2021, 5, 24), date(2022, 6, 6), date(2023, 5, 29), date(2024, 5, 20),
		date(2025, 6, 9), date(2026, 5, 25), date(2027, 5, 17), date(2028, 6, 5), date(2029, 5, 21),
		date(2030, 6, 10), date(2031, 6, 2), date(2032, 5, 17), date(2033, 6, 6),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsWhitMonday, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsWhitMonday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsWhitMonday, name)
}

func TestIsTrinitySunday(t *testing.T) {
	const name = "TrinitySunday"

	// Taken from https://en.wikipedia.org/wiki/Trinity_Sunday.
	tests := []time.Time{
		date(2014, 6, 15), date(2015, 5, 31), date(2016, 5, 22), date(2017, 6, 11), date(2018, 5, 27),
		date(2019, 6, 16), date(2020, 6, 7), date(2021, 5, 30), date(2022, 6, 12), date(2023, 6, 4),
		date(2024, 5, 26), date(2025, 6, 15), date(2026, 5, 31), date(2027, 5, 23), date(2028, 6, 11),
		date(2029, 5, 27), date(2030, 6, 16),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsTrinitySunday, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsTrinitySunday, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsTrinitySunday, name)
}

func TestIsCorpusChristi(t *testing.T) {
	const name = "CorpusChristi"

	// Taken from https://en.wikipedia.org/wiki/Corpus_Christi_(feast).
	tests := []time.Time{
		date(2013, 5, 30), date(2014, 6, 19), date(2015, 6, 4), date(2016, 5, 26), date(2017, 6, 15),
		date(2018, 5, 31), date(2019, 6, 20), date(2020, 6, 11), date(2021, 6, 3), date(2022, 6, 16),
		date(2023, 6, 8), date(2024, 5, 30), date(2025, 6, 19), date(2026, 6, 4), date(2027, 5, 27),
		date(2028, 6, 15), date(2029, 5, 31), date(2030, 6, 20),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsCorpusChristi, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsCorpusChristi, name)
	verifyYearErrorIsHoliday(t, easterSundayPerYearFirstYear-1, IsCorpusChristi, name)
}

func TestIsChristmasDay(t *testing.T) {
	const name = "ChristmasDay"

	tests := []time.Time{
		date(1959, 12, 25), date(1963, 12, 25), date(1990, 12, 25),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsChristmasDay, name)
	verifyRangeIsNotHoliday(t, tests, 0, IsChristmasDay, name)
}

func TestIsBoxingDay(t *testing.T) {
	tests := []time.Time{
		date(1959, 12, 26), date(1963, 12, 26), date(1990, 12, 26),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsBoxingDay, "BoxingDay")
}

func TestIsNewYearDay(t *testing.T) {
	tests := []time.Time{
		date(1959, 1, 1), date(1963, 1, 1), date(1990, 1, 1),
	}

	t.Parallel()
	verifyRangeIsHoliday(t, tests, 0, IsNewYearDay, "NewYearDay")
}

func date(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, &time.Location{})
}

func differ(d1, d2 time.Time) bool {
	return d1.Month() != d2.Month() || d1.Day() != d2.Day()
}

func verifyRange(t *testing.T, ts []time.Time, f func(int) (time.Time, error), s string) {
	t.Helper()

	for _, tt := range ts {
		y := tt.Year()

		act, err := f(y)
		if err != nil {
			t.Errorf("%v(%v): expected success '%v', got error %v", s, y, tt, err)

			return
		}

		if differ(tt, act) {
			t.Errorf("%v(%v): expected '%v', actual '%v'", s, y, tt, act)
		}
	}
}

func verifyYearDayRange(t *testing.T, ts []time.Time, f func(int) (int, error), s string) {
	t.Helper()

	for _, tt := range ts {
		y := tt.Year()

		act, err := f(y)
		if err != nil {
			t.Errorf("%v(%v): expected success '%v', got error %v", s, y, tt, err)

			return
		}

		if act != tt.YearDay() {
			t.Errorf("%v(%v): expected '%v', actual '%v'", s, y, tt.YearDay(), act)
		}
	}
}

func verifyYearRange(t *testing.T, y1 int, y2 int, f func(int) (time.Time, error), s string) {
	t.Helper()

	for y := y1; y <= y2; y++ {
		act, err := f(y)
		if err != nil {
			t.Errorf("%v(%v): expected success, got error %v", s, y, err)

			return
		}

		if act.Weekday() != time.Sunday {
			t.Errorf("%v(%v): expected 'Sunday', actual '%v'", s, y, act)
		}
	}
}

func verifyYearDayError(t *testing.T, f func(int) (int, error), y int, s string) {
	t.Helper()

	if _, err := f(y); err == nil {
		t.Errorf("%v(%v): xpected error, got success", s, y)
	}
}

func verifyYearError(t *testing.T, f func(int) (time.Time, error), y int, s string) {
	t.Helper()

	if _, err := f(y); err == nil {
		t.Errorf("%v(%v): xpected error, got success", s, y)
	}
}

func verifyRangeIsHoliday(t *testing.T, ts []time.Time, n int, f func(time.Time) (bool, error), s string) {
	t.Helper()

	for _, tt := range ts {
		tt = tt.AddDate(0, 0, n)

		act, err := f(tt)
		if err != nil {
			t.Errorf("Is%v('%v'): expected true, got error %v", s, tt, err)

			continue
		}

		if !act {
			t.Errorf("Is%v('%v'): expected true, actual false", s, tt)
		}
	}
}

func verifyRangeIsNotHoliday(t *testing.T, ts []time.Time, n int, f func(time.Time) (bool, error), s string) {
	t.Helper()

	for _, tt := range ts {
		tt = tt.AddDate(0, 0, n+1)

		act, err := f(tt)
		if err != nil {
			t.Errorf("Is%v('%v'): expected true, got error %v", s, tt, err)

			continue
		}

		if act {
			t.Errorf("Is%v('%v'): expected false, actual true", s, tt)
		}
	}
}

func verifyYearErrorIsHoliday(t *testing.T, y int, f func(time.Time) (bool, error), s string) {
	t.Helper()

	if _, err := f(date(y, 1, 1)); err == nil {
		t.Errorf("Is%v(%v): xpected error, got success", s, y)
	}
}
