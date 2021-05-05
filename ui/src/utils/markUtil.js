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
  for (let i = 0; i < editor.childNodes.length; i++) {
    const item = getParentSpanNodeIfNeeded(editor.childNodes[i])

    if (i < start || i > end) {
      const span = textToSpan(item)
      span.setAttribute('id', (k++).toString())
      allSlots.push(span)
    } else if (i === start) {
      const startLeft = slt.anchorOffset
      const startRight = isSame ? slt.focusOffset : startContainer.textContent.length
      const leftSection = startText.substr(0, startLeft)
      const rightSection = startText.substr(startLeft, startRight)

      // create part1 as span
      if (leftSection.length > 0) {
        const span = genSpan(leftSection, item)
        span.setAttribute('id', (k++).toString())
        allSlots.push(span)
      }
      // put part2 to cache
      if (isSame) {
        selectedIndex = k
        selectedText += startText.substr(startLeft, startRight - startLeft)
        const span1 = genSpan(selectedText, item)
        span1.setAttribute('id', (k++).toString())
        allSlots.push(span1)

        const span2 = genSpan(startText.substr(startRight), item)
        span2.setAttribute('id', (k++).toString())
        allSlots.push(span2)
      } else {
        selectedText += rightSection
      }
    } else if (i > start && i < end) {
      // put to cache
      selectedText += item.textContent
    } else if (i === end && !isSame) { // already be done if selection in single node
      const endLeft = 0
      const endRight = slt.focusOffset
      const leftSection = endText.substr(0, endLeft)
      const rightSection = endText.substr(endLeft, endRight - endLeft)

      // put part1 to cache
      selectedText += leftSection
      // create part2 as span
      if (rightSection.length > 0) {
        const span = genSpan(rightSection, item)
        span.setAttribute('id', (k++).toString())
        allSlots.push(span)
      }
    }

    if (i === end && !isSame) {
      selectedIndex = k

      const span = genSpan(selectedText)
      span.setAttribute('id', (k++).toString())
      allSlots.push(span)
    }
  }

  const mp = { allSlots: allSlots, selectedIndex: selectedIndex }
  console.log(mp)
  return mp
}

export const genSent = function (sentItems, slot) {
  const arr = []

  sentItems.forEach((item, index) => {
    const section = document.createElement('span')
    section.setAttribute('data-seq', index)
    section.setAttribute('data-type', slot.slotType)
    section.setAttribute('data-value', slot.value)
    section.innerText = item.text

    arr.push(section.outerHTML)
  })

  const html = arr.join('')
  return html
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
    span.setAttribute('data-type', 'synonym')
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
