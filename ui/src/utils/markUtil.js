export const convertSelectedToSlots = function (target, editor) {
  const allSlots = []
  let selectedIndex = 0

  target = getParentSpanNodeIfNeeded(target)
  console.log('---', target)

  const slt = window.getSelection()
  if (slt.toString() === '' || slt.rangeCount !== 1) return {}

  console.log('anchorNode', slt.anchorNode)
  console.log('anchorOffset', slt.anchorOffset)

  console.log('focusNode', slt.focusNode)
  console.log('focusOffset', slt.focusOffset)

  const range = slt.getRangeAt(0)
  console.log('range', range, range.toString())

  const startContainer = getParentSpanNodeIfNeeded(range.startContainer)
  const endContainer = getParentSpanNodeIfNeeded(range.endContainer)
  console.log('startContainer', startContainer, range.startOffset)
  console.log('endContainer', endContainer, range.endOffset)
  const isSame = startContainer === endContainer
  console.log('is same', isSame)

  let start = -1
  let end = -1
  for (let i = 0; i < editor.childNodes.length; i++) {
    const item = getParentSpanNodeIfNeeded(editor.childNodes[i])

    if (item === startContainer) {
      console.log('start')
      start = i
    }
    if (item === endContainer) {
      console.log('end')
      end = i
    }
  }

  const startText = startContainer.textContent
  const endText = endContainer.textContent

  let selectedText = ''
  let k = 0
  for (let j = 0; j < editor.childNodes.length; j++) {
    const item = getParentSpanNodeIfNeeded(editor.childNodes[j])

    if (j < start || j > end) {
      const span = textToSpan(item)
      span.setAttribute('id', (k++).toString())
      allSlots.push(span)
    } else if (j === start) {
      const startLeft1 = slt.anchorOffset
      const startRight1 = isSame ? slt.focusOffset : startContainer.textContent.length
      const leftSection1 = startText.substr(0, startLeft1)
      const rightSection1 = startText.substr(startLeft1, startRight1 - startLeft1)

      // create part1 as span
      if (leftSection1.length > 0) {
        const span = genSpan(leftSection1, item)
        span.setAttribute('id', (k++).toString())
        allSlots.push(span)
      }
      // put part2 to cache
      if (isSame) {
        selectedIndex = k
        selectedText += startText.substr(startLeft1, startRight1 - startLeft1)
        const span1 = genSpan(selectedText, item)
        span1.setAttribute('id', (k++).toString())
        allSlots.push(span1)

        const span2 = genSpan(startText.substr(startRight1), item)
        span2.setAttribute('id', (k++).toString())
        allSlots.push(span2)
      } else {
        selectedText += rightSection1
      }
    } else if (j > start && j < end) {
      // put to cache
      selectedText += item.textContent
    } else if (j === end && !isSame) { // already be done if selection in single node
      const endLeft2 = 0
      const endRight2 = slt.focusOffset
      const leftSection2 = endText.substr(0, endLeft2 + 1)
      const rightSection2 = endText.substr(endLeft2 + 1, endRight2 - endLeft2 - 1)

      // put part1 to cache
      selectedText += leftSection2
      // create part2 as span
      if (rightSection2.length > 0) {
        const span = genSpan(rightSection2, item)
        span.setAttribute('id', (k++).toString())
        allSlots.push(span)
      }
    }

    if (j === end && !isSame) {
      selectedIndex = k

      const span = genSpan(selectedText, item)
      span.setAttribute('id', (k++).toString())
      allSlots.push(span)
    }
  }

  const mp = { allSlots: allSlots, selectedIndex: selectedIndex }
  console.log(mp)
  return mp
}

export const genSent = function (allSlots, selectedIndex, slot) {
  const arr = []
  console.log('---', selectedIndex)

  allSlots.forEach((item, index) => {
    const section = document.createElement('span')
    let dataType = item.getAttribute('data-type')
    let dataValue = item.getAttribute('data-value')

    if (index === selectedIndex) {
      dataType = slot.slotType
      dataValue = slot.value
    }

    section.setAttribute('id', index)
    if (dataType) {
      section.setAttribute('data-type', dataType)
      addCls(section, dataType)
    }

    if (dataValue) section.setAttribute('data-value', dataValue.toString())
    else section.setAttribute('data-value', '')
    section.innerText = item.innerText

    arr.push(section.outerHTML)
  })

  return arr.join('')
}

export const genSentSlots = function (editor) {
  const slots = []
  editor.childNodes.forEach((item, index) => {
    const slotObj = {}

    slotObj.seq = index
    slotObj.id = item.id
    slotObj.type = item.getAttribute('data-type')
    slotObj.value = item.getAttribute('data-value')
    slotObj.text = item.innerText

    slots.push(slotObj)
  })

  return slots
}

export const getParentSpanNodeIfNeeded = function (target) {
  if (target.parentNode && target.parentNode.nodeName.toLowerCase() === 'span') {
    target = target.parentNode
  }
  return target
}

export const genSpan = function (text, node) {
  const span = document.createElement('span')
  span.innerText = text

  if (node.getAttribute) {
    span.setAttribute('data-type', node.getAttribute('data-type'))
    span.setAttribute('data-value', node.getAttribute('data-value'))
  } else {
    // span.setAttribute('data-type', 'synonym')
  }

  return span
}

export const textToSpan = function (node) {
  if (node.nodeName === '#text') {
    const span = document.createElement('span')
    span.innerText = node.nodeValue
    return span
  } else {
    return node
  }
}

export const addCls = function (element, value) {
  if (!element.className) {
    element.className = value
  } else {
    let newClassName = element.className
    newClassName += ' '
    newClassName += value
    element.className = newClassName
  }
}
