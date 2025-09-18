<?php

// Fungsi untuk membuat sitemap.xml
function generateSitemap($list, $hostname, $sitemapFileName) {
    $sitemapContent = '<?xml version="1.0" encoding="UTF-8"?>' . PHP_EOL;
    $sitemapContent .= '<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">' . PHP_EOL;

    foreach ($list as $url) {
        $url = trim($url); // Menghilangkan spasi di awal dan akhir
        $url = str_replace(' ', '+', $url); // Mengganti spasi di tengah dengan +

        // Membuat URL dengan ?news_stecu=
        $fullUrl = $hostname . '?news_stecu=' . $url;

        $sitemapContent .= '  <url>' . PHP_EOL;
        $sitemapContent .= '    <loc>' . htmlspecialchars($fullUrl) . '</loc>' . PHP_EOL;
        $sitemapContent .= '  </url>' . PHP_EOL;
    }

    $sitemapContent .= '</urlset>' . PHP_EOL;

    file_put_contents($sitemapFileName, $sitemapContent);
    echo "Sitemap berhasil dibuat: $sitemapFileName" . PHP_EOL;
}

// Fungsi untuk membuat robots.txt
function generateRobotsTxt($hostname, $sitemapFileName) {
    $robotsContent = "User-agent: *" . PHP_EOL;
    $robotsContent .= "Disallow:" . PHP_EOL;
    $robotsContent .= "Sitemap: " . htmlspecialchars($hostname . $sitemapFileName) . PHP_EOL; // Menggunakan URL tanpa ?news_stecu=

    file_put_contents('robots.txt', $robotsContent);
    echo "Robots.txt berhasil dibuat." . PHP_EOL;
}

// Baca list.txt
$list = file('list.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);

if (!$list) {
    echo "Gagal membaca list.txt." . PHP_EOL;
    exit;
}

// Tentukan hostname dan nama file sitemap
$hostname = 'https://www.tomonet.co.id/'; // Pastikan ini hanya hostname
$sitemapFileName = 'ga.xml'; // Nama file sitemap

// Buat sitemap.xml dan robots.txt
generateSitemap($list, $hostname, $sitemapFileName);
generateRobotsTxt($hostname, $sitemapFileName);

?>
