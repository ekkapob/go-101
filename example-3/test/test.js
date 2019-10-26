const puppeteer = require('puppeteer');
const assert = require('assert');

var browser, page;

before(async () => {
  browser = await puppeteer.launch();
  page = await browser.newPage();
});

after(async () => {
  browser.close();
});

const tests = [
  {
    route: '/',
    expectedContent: 'Hello World',
  },
  {
    route: '/log',
    expectedContent: 'Hello World with Middlewares',
  },
  {
    route: '/header',
    expectedContent: 'Hello World with Middlewares',
  },
  {
    route: '/authen/pass',
    expectedContent: 'Congrats!',
  },
  {
    route: '/authen/fail',
    expectedContent: 'Hello World',
  },
  {
    route: '/not-found',
    expectedContent: 'Hello World',
  },
];

tests.forEach((test) => {
  describe(`at '${test.route}'`, () => {
    it(`should show ${test.expectedContent}`, async () => {
      await page.goto(`${process.env.HOST}${test.route}`);
      let content = await page.content();
      assert.equal(RegExp(test.expectedContent).test(content), true);
    });
  });
});
