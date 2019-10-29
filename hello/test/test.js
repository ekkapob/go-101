const puppeteer = require('puppeteer');
const assert = require('assert');

var browser, page;

before(async () => {
  // browser = await puppeteer.launch({headless: false});
  browser = await puppeteer.launch();
  page = await browser.newPage();
});

after(async () => {
  browser.close();
});

describe('greeting', () => {
  it('should show proper greeting', async () => {
    await page.goto(`${process.env.HOST}/`);
    let content = await page.content();
    await page.type('input[type="text"]', 'Justin');
    // await sleep(2000);
    await page.click('button[type="submit"]');
    // await sleep(2000);
    const greeting = await page.$eval('h1', (element) => {
      return element.innerHTML;
    });
    assert.equal(greeting, "Hello Justin");
  });
});

describe('back to home page', () => {
  it('should return to homepage when click home link', async () => {
    await page.goto(`${process.env.HOST}/`);
    let content = await page.content();
    await page.type('input[type="text"]', 'Justin');
    await page.click('button[type="submit"]');
    // await sleep(2000);
    await page.waitForSelector('a');
    await page.click('a');
    // await sleep(2000);
    const label = await page.$eval('label', (element) => {
      return element.innerHTML;
    });
    assert.equal(RegExp(/Type your name/).test(label), true);
  });
});

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}
