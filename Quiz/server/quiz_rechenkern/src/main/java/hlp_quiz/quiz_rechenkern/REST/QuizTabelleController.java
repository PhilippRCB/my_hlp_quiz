package hlp_quiz.quiz_rechenkern.REST;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import com.fasterxml.jackson.annotation.JacksonInject;

import hlp_quiz.quiz_rechenkern.model.FrageBoard;

@RestController
public class QuizTabelleController {

    private final FrageBoard frageBoard;
    
    public QuizTabelleController(@JacksonInject FrageBoard frageBoard) {
        this.frageBoard = frageBoard;
    }

    @GetMapping("/quizTabelle")
    public FrageBoard frageBoard()  {return frageBoard; }
}