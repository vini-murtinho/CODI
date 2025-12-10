const puppeteer = require('puppeteer');
const path = require('path');
const fs = require('fs');

(async () => {
  try {
    const browser = await puppeteer.launch();
    const page = await browser.newPage();
    
    // Caminho absoluto do arquivo HTML
    const htmlPath = path.join(__dirname, 'APRESENTACAO.html');
    const fileUrl = `file://${htmlPath}`;
    
    await page.goto(fileUrl, { waitUntil: 'networkidle0' });
    
    // Configurar PDF com opções de qualidade
    await page.pdf({
      path: path.join(__dirname, 'APRESENTACAO.pdf'),
      format: 'A4',
      margin: {
        top: '20px',
        right: '20px',
        bottom: '20px',
        left: '20px'
      },
      printBackground: true,
      landscape: false
    });
    
    console.log('✅ PDF exportado com sucesso: APRESENTACAO.pdf');
    
    await browser.close();
    process.exit(0);
  } catch (error) {
    console.error('❌ Erro ao exportar PDF:', error);
    process.exit(1);
  }
})();
