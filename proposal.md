# Project Proposal: Pwnd-Buddy

![alt text](arthur-mazi-a8CxRWIu8yw-unsplash.jpg)

## Introduction

In today's digital age, concerns about online privacy and data security are paramount. With numerous data breaches occurring daily, individuals and organizations alike are at risk of having their sensitive information compromised. To help maintaining assets secured, I propose to develop a tool that will keep owners aprised of wheter or not an account has been compromised.

## Problem Statement

Many individuals use their email addresses across multiple platforms and services, making them vulnerable to potential data breaches. However, staying informed about whether one's email has been part of a breach is a challenging task. Current solutions often require manual searches on platforms with incessant adds.

## Proposed Solution

The Pwnd-Buddy aims to provide a quick and easy way for users to determine if their email addresses have been involved in any known data breaches. By leveraging several APIs and the implementation of simple CLI commads, Pwnd-Buddy will offer a centralized platform for users to query the status of their email addresses.

## Features

1. **Email Address Input**: Users can input one or more email addresses they wish to check for breaches.
2. **Breach Detection**: The tool will query trusted breach data sources and APIs to identify if the provided email addresses have been part of any known breaches.
3. **Notification System**: Users will receive immediate notifications if any of their email addresses are found in a breach, along with relevant details.

## Technology Stack

- **Programming Language**: Go (Golang) will be utilized for its concurrency features and efficiency in handling HTTP requests.
- **API Integration**: The tool will integrate with the HaveIBeenPwnd API.
- **Data Storage**: Stateless application, Pwnd-Buddy does not store your data. Everything is displayed in the CLI. You do however have the option to save your results to a file in the root. 

## Target Audience

The Email Breach Checker Tool will cater to individuals who are concerned about their online security and privacy, including:

- General internet users
- IT professionals
- Security-conscious individuals and organizations

## Initial Implementation Plan

1. **Research and Requirements Gathering**: Identify APIs to interact with and various OSINT tools that could be useful.
2. **Prototype Development**: Develop a minimal viable product (MVP) to demonstrate core functionality, including email address input and breach detection.
3. **API Integration**: Integrate with selected API.
4. **User Interface Design (CLI)**: Design a user-friendly prompts and commands for inputting email addresses and viewing results
5. **Testing**: Conduct extensive testing to ensure accuracy and reliability.


## Conclusion

Pwnd-Buddy aims to empower users with knowledge and awareness regarding potential data breaches involving their email addresses. By providing a streamlined and proactive solution, the tool seeks to mitigate the risks associated with online security threats, ultimately fostering a safer digital environment for all.
