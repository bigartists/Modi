import styled from 'styled-components'
import { SPACE_TIMES } from 'assets/styles/styledcom/StyleConstants'
import bgVedio from 'assets/vedio/bg.mp4'
import { useEffect, useRef } from 'react'
interface ITopic {
  topic: string
  description: string
}

function Topic({ topic, description }: ITopic) {
  const vedioRef = useRef<HTMLVideoElement>(null)

  useEffect(() => {
    let interval: NodeJS.Timeout
    let vedio = vedioRef?.current

    interval = setTimeout(() => {
      vedio && vedio.play()
    }, 1000)

    return () => {
      vedio && vedio.pause()
      clearTimeout(interval)
    }
  }, [])
  return (
    <Wrapper className="topic">
      <div className="panel">
        <div className="category">{topic}</div>
        <div className="desc">{description}</div>
      </div>
      <div className="vedioPanel">
        <video
          className="video"
          loop
          muted
          src={bgVedio}
          ref={vedioRef}
          // style={{ height: '420px' }}
        ></video>
      </div>
    </Wrapper>
  )
}

const Wrapper = styled.article`
  height: ${SPACE_TIMES(64)};
  mix-blend-mode: normal;
  display: flex;
  flex-direction: column;
  /* background: linear-gradient(180deg, #f0f5fb, #ffffff); */
  background-size: 800px 100%;
  position: relative;
  .panel {
    position: absolute;
    z-index: 2;
    width: 100%;
    height: 100%;
    padding-top: ${SPACE_TIMES(17)};
    padding-left: ${SPACE_TIMES(18)};
    .category {
      font-size: ${SPACE_TIMES(6)};
      font-weight: 700;
    }
    .desc {
      font-size: 14px;
      margin-top: ${SPACE_TIMES(3)};
      width: 476px;
    }
  }
  .vedioPanel {
    position: absolute;
    width: 100%;
    height: 100%;
    z-index: 1;
    overflow: hidden;
    .video {
      margin: 0px;
      object-fit: cover;
      padding: 0px;
      position: absolute;
      width: 100%;
    }
  }
`

export default Topic
