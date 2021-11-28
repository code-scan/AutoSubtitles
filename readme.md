# AutoSubtitles 自动给视频生成中文字幕并合成

## 下载

[Releases](https://github.com/code-scan/AutoSubtitles/releases)

## 配置

需要修改`config.default.ini`为你自己的配置，并将其重命名为`config.ini`


## 运行逻辑

首先使用ffmpeg提取声轨，然后使用**录音文件识别**提取出对应的文字，再使用**通用翻译**进行翻译，最后合成srt字幕，最后再使用ffmpeg将srt和视频合并。

## 依赖

需要开通阿里云**通用翻译**与**录音文件识别**，并且需要在其中新建项目，配置文件中的**appkey**为**录音文件识别**的项目appkey


https://www.aliyun.com/product/ai/base_alimt?spm=5176.21213303.1378385.1.49063eda7XMipW&scm=20140722.S_card@@%E5%8D%A1%E7%89%87@@784._.ID_card@@%E5%8D%A1%E7%89%87@@784-RL_%E7%BF%BB%E8%AF%91-OR_ser-V_2-P0_0

https://ai.aliyun.com/nls/filetrans?spm=5176.12061031.1228726.1.47fe3cb43I34mn