Steps for Linux and Mac

To setup the chromedriver, following the given below steps:
1. Install chrome on your laptop and check out its version.

    ```bash
    google-chrome --version 
    ```
2. Now, depending on the chrome version, download a web-driver.

    - If you are using Chrome version 101, please download ChromeDriver 101.0.4951.15
    - If you are using Chrome version 100, please download ChromeDriver 100.0.4896.60 (I am currently using this)
    - If you are using Chrome version 99, please download ChromeDriver 99.0.4844.51


    ```bash
    wget https://chromedriver.storage.googleapis.com/94.0.4606.61/chromedriver_linux64.zip && unzip chromedriver_linux64.zip && sudo mv chromedriver /usr/bin/chromedriver && sudo chown root:root /usr/bin/chromedriver && sudo chmod +x /usr/bin/chromedriver 
    ```
3. export ACK_GINKGO_RC=true